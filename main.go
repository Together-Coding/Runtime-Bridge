package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/together-coding/runtime-bridge/containers"
	"github.com/together-coding/runtime-bridge/db"
	"github.com/together-coding/runtime-bridge/docs"
	"github.com/together-coding/runtime-bridge/runtimes"
	"github.com/together-coding/runtime-bridge/users"
	"github.com/together-coding/runtime-bridge/utils"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	// Load configs
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	// Initialize DB
	db.Initialize()

	// Gin
	r := gin.Default()
	utils.CORSMiddleware(r)

	// users
	userApi := r.Group("/api/users")
	userApi.Use(users.IdentifyUser())
	users.Register(userApi.Group("/"))

	// runtimes
	runtimeApi := r.Group("/api/runtimes")
	runtimes.Register(runtimeApi.Group("/"))

	// containers
	containerApi := r.Group("/api/containers")
	containers.Register(containerApi.Group("/"), users.IdentifyUser)

	// swagger
	if utils.GetConfigBool("DEBUG") {
		docs.SwaggerInfo.BasePath = "/api"
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	r.Run()
}
