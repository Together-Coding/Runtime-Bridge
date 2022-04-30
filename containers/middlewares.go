package containers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifyApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check X-API-KEY to see whether it is from Runtime Agent
		apiKey := c.Request.Header.Get("X-API-KEY")
		if apiKey != "" {
			// X-API-KEY is valid
			alloc := GetContainerFromKey(apiKey)
			if alloc.isAvailable() {
				c.Set("alloc", *alloc)
				c.Next()
			} else {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "No container found."})
			}
		} else {
			// Unauthorized request
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "X-API-KEY is missing."})
		}
	}
}
