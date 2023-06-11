package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	mysql "pasarmalamkebaya/database/init"
)

func main() {

	// Install Gin!
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Install air -> live reload for go apps -> can just run "air" now!

	// Install Migrate -> Makefile!

	// Install Gorm

	// Install Godotenv

	// Initializing Database Connection
	db := mysql.GetDatabaseConnection()

	r.Run() // listen and serve on localhost:8080

}
