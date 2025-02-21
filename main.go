package main

import (
	"log"
	"net/http"
	"os"

	"github.com/JavadsGithub/goath/auth"
	"github.com/JavadsGithub/goath/auth/validators"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	r := gin.Default()

	r.POST("/authenticate", func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		validatorKey := os.Getenv("JWT_VALIDATOR")
		if len(validatorKey) < 1 {
			validatorKey = "default"
		}
		validator := validators.GetValidators()[validatorKey]
		if validator == nil {
			validator = validators.DefaultClaimsValidator
		}

		isAuthenticated, err := auth.Authenticate(token, validator)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
		}

		if isAuthenticated {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"msg":     "Authorized",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Unauthorized",
			})
		}

	})

	r.Run()
}
