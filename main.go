package main

import (
	"github.com/gin-gonic/gin"
	"rest-api-with-golang/controllers"
	"rest-api-with-golang/initializers"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/api/v1/posts", controllers.PostsCreate)
	r.GET("/api/v1/posts", controllers.PostsIndex)
	r.GET("/api/v1/posts/:postId", controllers.PostsDetail)
	r.PUT("/api/v1/posts/:postId", controllers.PostsUpdate)
	r.DELETE("/api/v1/posts/:postId", controllers.PostsDelete)

	r.Run()
}
