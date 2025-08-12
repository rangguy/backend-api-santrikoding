package main

import (
	"github.com/gin-gonic/gin"
	"rangguy/backend-api/config"
	"rangguy/backend-api/database"
)

func main() {
	config.LoadEnv()

	database.InitDB()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
