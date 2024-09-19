package main

import (
	"hackathon/backend/controllers"
	"hackathon/backend/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnecToDB()
}

func main() {

	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/postsIndex", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.Run()
}
