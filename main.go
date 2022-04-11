package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/together-coding/runtime-bridge/db"
)

func main() {
	// Load configs
	dotenvErr := godotenv.Load(".env")
	if dotenvErr != nil {
		panic(dotenvErr)
	}

	// Initialize DB
	db.Initialize()

	// Gin
	r := gin.Default()

	r.Run()
}
