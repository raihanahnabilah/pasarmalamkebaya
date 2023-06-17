package main

import (
	"github.com/gin-gonic/gin"

	mysql "pasarmalamkebaya/database/init"
	handler "pasarmalamkebaya/handler/http"
	"pasarmalamkebaya/repository"
	"pasarmalamkebaya/usecase"
)

func main() {

	// Install Gin!
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	router := gin.Default()

	// Install air -> live reload for go apps -> can just run "air" now!

	// Install Migrate -> Makefile!

	// Install Gorm

	// Install Godotenv

	// Initializing Database Connection
	db := mysql.GetDatabaseConnection()

	// To plug in repo, usecase, and handler
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	registerUsecase := usecase.NewRegisterUsecase(userRepo, userUsecase)
	registerHandler := handler.NewRegisterHandler(registerUsecase)
	registerHandler.Route(&router.RouterGroup)

	// API Routers!
	// api := router.Group("api/v1")
	// api.GET("/register", registerHandler.Register)

	router.Run() // listen and serve on localhost:8080

}
