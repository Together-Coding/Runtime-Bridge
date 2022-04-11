package containers

import "github.com/gin-gonic/gin"

func Register(router *gin.RouterGroup) {
	router.POST("/launch", LaunchContainer)
}

// LaunchContainer is called when the user tries to connect or to use
// runtime-related functionalities. This endpoint launches new Docker Container
// that is allocated for the user and return its information back after it is
// launched.
func LaunchContainer(c *gin.Context) {
	// Check whether the user has already requested, and new container is
	// in launch or in active status. If it is, return its information right away.

	// Launch a specified container

	// Ping the container

	// Save container's information into Database

	// Return the data to the user

	// And **start monitoring**
}
