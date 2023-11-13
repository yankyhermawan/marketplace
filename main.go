package main

import (
	"github.com/gin-gonic/gin"

	"github.com/yankyhermawan/marketplace/database"
	"github.com/yankyhermawan/marketplace/interfaces"
	"github.com/yankyhermawan/marketplace/registerService"
)

func main() {
	db := database.InitDB()
	database.MigrateDB(db)
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	router.POST("/register", func(c *gin.Context) {
		var requestBody interfaces.RequestBody

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		response := registerService.RegisterUser(&requestBody)
		c.JSON(response.Code, gin.H{"message": response.Response})

	})

	router.Run(":4000")
}
