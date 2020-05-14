package main

import (
	Utils "github.com/electra-systems/athena/utils"

	"github.com/gin-gonic/gin"
)

func Server() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Homepage",
		})
	})

	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey the",
		})
	})

	r.Run()
}

func main() {
	Utils.RunTests()
	Utils.TestRedis()
}
