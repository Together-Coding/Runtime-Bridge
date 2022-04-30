package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func IdentifyUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		authPrefix := "Bearer "

		if strings.HasPrefix(auth, authPrefix) {
			// Parse auth token
			idx := strings.Index(auth, authPrefix) + len(authPrefix)
			user := VerifyUser(c, auth[idx:])
			if !user.Valid {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": user.ErrorReason})
			}
			c.Set("user", user)
			c.Next()
		} else {
			// Unauthorized request
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized"})
		}
	}
}
