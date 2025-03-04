package main

import (
	"github.com/JavadsGithub/goath/config"
	"github.com/JavadsGithub/goath/controllers"
	"github.com/JavadsGithub/goath/repositories"
	"github.com/JavadsGithub/goath/services"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(*userRepo)
	userController := controllers.NewUserController(*userService)

	router := gin.Default()

	router.GET("/users/:id", userController.GetUserById)
	router.POST("/users", userController.CreateUser)

	router.Run(":8080")
}
