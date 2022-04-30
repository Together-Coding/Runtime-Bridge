package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(router *gin.RouterGroup) {
	router.GET("/", ApiIndex)
}

type IndexResp struct {
	Email string `json:"email" example:"usera2tt@gmail.com"`
}

// ApiIndex is a test page
func ApiIndex(c *gin.Context) {
	user, _ := c.Get("user")

	res := IndexResp{Email: user.(VerifiedUser).Email}
	c.JSON(http.StatusOK, res)
}
