package main

import (
	"Posts-Demo/controllers"
	"Posts-Demo/inits"

	"github.com/gin-gonic/gin"
)

func init() {

	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.POST("/post", controllers.CreatePost)

	r.GET("/posts", controllers.GetPosts)

	r.GET("/post/:id", controllers.GetPost)

	r.PUT("/post/:id", controllers.UpdatePost)

	r.DELETE("/post/:id", controllers.DeletePost)

	r.Run()
}
