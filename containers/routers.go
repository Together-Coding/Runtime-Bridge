package containers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/together-coding/runtime-bridge/aws"
	"github.com/together-coding/runtime-bridge/db"
	"github.com/together-coding/runtime-bridge/runtimes"
	"github.com/together-coding/runtime-bridge/users"
	"github.com/together-coding/runtime-bridge/utils"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Register(router *gin.RouterGroup, mdwIdentifyUser utils.MiddlewareType) {
	router.POST("/launch", mdwIdentifyUser(), LaunchContainer)
	router.GET("/info", VerifyApiKey(), ContainerInfo)
}

type LaunchReq struct {
	RuntimeImageID int64 `json:"image_id"`
}

type LaunchResp struct {
	Url  *string `json:"url" example:"127.0.0.1"`
	Port *uint16 `json:"port" example:"8000"`
}

// LaunchContainer godoc
// @Summary Launch container if not assigned, and then return the assigned one.
// @Description This is called when the user tries to connect or to use runtime-related functionalities. <br>This endpoint launches new Docker Container that will be allocated for the user if not allocated yet, and then returns its information in order for the user to access it.
// @Tags containers
// @Accept json
// @param Authorization header string true "User's Json Web Token" example(Bearer ...)
// @Param name body string true "target runtime name" example(C gcc11)
// @Param lang body string true "target language" example(C)
// @Produce json
// @Success 200 {object} LaunchResp
// @Failure 400 {object} utils.CommonErrorResp "'`name` or `lang` is missing.' | 'It is launching. Please wait for a moment.'"
// @Router /containers/launch [post]
func LaunchContainer(c *gin.Context) {
	user_, _ := c.Get("user")
	body := LaunchReq{}
	err := c.BindJSON(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.CommonErrorResp{
			Msg: "`name` or `lang` is missing.",
		})
		return
	}
	user := user_.(users.VerifiedUser)

	// Check database to see whether the user has already requested, and new container is
	// in launching or active status. If it is, return it right away.
	// TODO: support allocating multiple containers for a user (with upperbound limit),
	//   but restrict to one per course.

	runtimeImage := runtimes.GetRuntimeImage(body.RuntimeImageID)
	alloc := RuntimeAllocation{}
	db.DB.Where(&RuntimeAllocation{
		UserID:         user.UserId,
		RuntimeImageID: runtimeImage.ID,
		Health:         int8(ContActive),
	}).Or(&RuntimeAllocation{
		UserID:         user.UserId,
		RuntimeImageID: runtimeImage.ID,
		Health:         int8(ContLaunching),
	}).Find(&alloc)
	fmt.Println(alloc.ID, alloc.UserID, alloc.RuntimeImageID, alloc.Health, alloc.CreatedAt)

	if alloc.Health == int8(ContActive) {
		// Return from already allocated
		c.AbortWithStatusJSON(http.StatusOK, LaunchResp{
			Url:  &alloc.ContIp,
			Port: &alloc.ContPort,
		})
		return
	} else if alloc.Health == int8(ContLaunching) {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.CommonErrorResp{
			Msg: "It is being launched. Please wait for a moment.",
		})
		return
	}

	// Create new allocation with "launching" status
	contAPIKey := utils.RandString(32, true, true, true)

	alloc = RuntimeAllocation{
		UserID:         user.UserId,
		UserIp:         c.ClientIP(),
		ContAPIKey:     contAPIKey,
		Health:         int8(ContLaunching),
		RuntimeImageID: runtimeImage.ID,
		//ContLaunchedAt: time.Now(),
		//ContIp:         *contIp,
		//ContPort:       uint16(contPort),
		//ContUser:       "together",
		//ContAuthType:   "password",
		//ContAuth:       "",
	}
	db.DB.Create(&alloc)

	// Launch a specified container
	clusterName := os.Getenv("CLUSTER_NAME")
	if utils.GetConfigBool("DEBUG") {
		clusterName += "-dev"
	}

	taskDef := fmt.Sprintf("%v:%v", runtimeImage.Taskdef, runtimeImage.Revision)
	startedBy := fmt.Sprintf("u%v", user.UserId)

	useLocalCont, _ := strconv.ParseBool(os.Getenv("LOCAL_CONTAINER"))
	var contIp string
	if !useLocalCont {
		ecsTask := aws.StartContainer(clusterName, taskDef, startedBy)
		aws.DescTasks(ecsTask)

		contIp = *ecsTask.PublicIp
		alloc.ContLaunchedAt = *ecsTask.Task.CreatedAt
	} else {
		contIp = "127.0.0.1"
		alloc.ContLaunchedAt = time.Now()
	}
	alloc.ContIp = contIp
	alloc.Health = int8(ContActive)
	db.DB.Save(alloc)

	// Ping the container until get response
	agentPort := utils.GetConfigUint16("AGENT_PORT")
	MustPingAgent(contIp, agentPort, 300)

	// Init the container
	// Send API key that will be used for communication between bridge and agent servers
	// and get ssh credentials
	initResp := InitAgent(contIp, agentPort, contAPIKey, 3)

	// Save container's information into Database
	alloc.ContUser = initResp["username"].(string)
	alloc.ContAuthType = initResp["auth_type"].(string)
	alloc.ContAuth = initResp["auth"].(string)
	alloc.ContPort = agentPort
	db.DB.Save(alloc)

	// Return data to the user
	c.JSON(http.StatusOK, LaunchResp{
		Url:  &contIp,
		Port: &agentPort,
	})

	// TODO:
	// Start to monitor this container.
	// The monitor thread should be already spawned (sync required)
	//  or should be executed in separate servers
}

type MyContResp struct {
	ContIp       string `json:"cont_ip"`
	ContPort     uint16 `json:"cont_port"`
	ContUser     string `json:"cont_user"`
	ContAuthType string `json:"cont_auth_type"`
	ContAuth     string `json:"cont_auth"`
}

// ContainerInfo godoc
// @Summary Return the containers' data matched with the API key
// @Description It returns the allocated container's data that is matched with the API key. <br>Normally, only Runtime Agent servers call this endpoint.<br> If `X-API-KEY` is missing, 400 error is returned.
// @Tags containers
// @Accept json
// @Produce json
// @param X-API-KEY header string true "API key assigned to each Agent server."
// @Success 200 {object} MyContResp
// @Failure 404 {object} utils.CommonErrorResp
// @Router /containers/info [get]
func ContainerInfo(c *gin.Context) {
	alloc_, _ := c.Get("alloc")
	alloc := alloc_.(RuntimeAllocation)

	if alloc.ContAuth == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.CommonErrorResp{
			Msg: "No container found.",
		})
		return
	}

	c.JSON(http.StatusOK, MyContResp{
		ContIp:       alloc.ContIp,
		ContPort:     alloc.ContPort,
		ContUser:     alloc.ContUser,
		ContAuthType: alloc.ContAuthType,
		ContAuth:     alloc.ContAuth,
	})
}
