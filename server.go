package main

import (
	"github.com/JavadsGithub/goath/config"
	"github.com/JavadsGithub/goath/controllers"
	"github.com/JavadsGithub/goath/middleware"
	"github.com/JavadsGithub/goath/repositories"
	"github.com/JavadsGithub/goath/services"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(*userRepo)
	userController := controllers.NewUserController(*userService)
	authController := controllers.NewAuthController(*userService)

	router := gin.Default()

	// Public routes
	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)

	// Protected routes
	authRoutes := router.Group("/").Use(middleware.AuthMiddleware())
	{
		authRoutes.GET("/users/:id", userController.GetUserById)
	}

	router.Run(":8080")
}
