package utils

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type MiddlewareType func() gin.HandlerFunc

func CORSMiddleware(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"https://together-coding.com",
		"https://*.together-coding.com",
		"http://localhost:3000",
		"http://127.0.0.1:3000",
	}
	config.AllowHeaders = append(config.AllowHeaders, []string{
		"Authorization",
		"X-API-KEY",
	}...)
	config.AllowWildcard = true

	router.Use(cors.New(config))
}
