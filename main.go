package main

import (
	"Posts-Demo/inits"

	"github.com/gin-gonic/gin"
)

func init() {

	inits.LoadEnv()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.Run()
}
