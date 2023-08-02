package main

import (
	"Posts-Demo/controllers"
	"Posts-Demo/inits"
	"Posts-Demo/middlewares"

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

	// Post Routes

	r.POST("/post", middlewares.RequireAuth, controllers.CreatePost)

	r.GET("/posts", controllers.GetPosts)

	r.GET("/post/:id", controllers.GetPost)

	r.PUT("/post/:id", controllers.UpdatePost)

	r.DELETE("/post/:id", controllers.DeletePost)

	// User Routes

	r.POST("/signup", controllers.Signup)

	r.POST("/login", controllers.Login)

	r.GET("/users", controllers.GetUsers)

	r.POST("/auth", controllers.Validate)

	r.GET("/logout", controllers.Logout)

	r.Run()
}
