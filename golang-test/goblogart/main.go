package main

import (
	"goblogart/controllers"
	"goblogart/inits"
	"goblogart/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnv()
	inits.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", controllers.GetPosts)
	r.GET("/:id", controllers.GetPost)
	r.POST("/", middlewares.RequireAuth, controllers.CreatePost)
	r.PUT("/:id", controllers.UpdatePost)
	r.DELETE("/:id", controllers.DeletePost)
	r.POST("/users", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/auth", controllers.Validate) // not working?
	r.GET("/users", controllers.GetUsers)
	r.GET("/logout", controllers.Logout)

	r.Run()
}
