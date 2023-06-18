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

	// For Registering
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	mailUsecase := usecase.NewMailUsecase()
	registerUsecase := usecase.NewRegisterUsecase(userRepo, userUsecase, mailUsecase)
	registerHandler := handler.NewRegisterHandler(registerUsecase)
	registerHandler.Route(&router.RouterGroup)

	// For Login!
	oauthAccessRepo := repository.NewOauthAccessRepository(db)
	oauthClientRepo := repository.NewOauthClientRepository(db)
	oauthUsecase := usecase.NewOauthUsecase(oauthClientRepo)
	loginUsecase := usecase.NewLoginUsecase(oauthAccessRepo, oauthUsecase, userUsecase)
	loginHandler := handler.NewLoginHandler(loginUsecase)
	loginHandler.Route(&router.RouterGroup)

	// API Routers!
	// api := router.Group("api/v1")
	// api.GET("/register", registerHandler.Register)

	router.Run() // listen and serve on localhost:8080
}
