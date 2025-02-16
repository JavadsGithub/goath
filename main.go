package main

import (
	"net/http"

	"github.com/JavadsGithub/goath/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/authenticate", func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		isAuthenticated, err := auth.Authenticate(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
		}

		if isAuthenticated {
			c.JSON(http.StatusOK, gin.H{
				"msg": "Authorized",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
		}

	})

	r.Run()
}
