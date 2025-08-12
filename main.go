package main

import (
	"github.com/gin-gonic/gin"
	"rangguy/backend-api/config"
)

func main() {
	config.LoadEnv()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
