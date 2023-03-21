package main

import (
	"github.com/gin-gonic/gin"
	"rest-api-with-golang/controllers"
	"rest-api-with-golang/initializers"
	"rest-api-with-golang/middleware"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/api/v1/auth/signup", controllers.SignUp)
	r.POST("/api/v1/auth/login", controllers.SignIn)

	// POST
	r.POST("/api/v1/posts", middleware.AuthMiddleware, controllers.PostsCreate)
	r.GET("/api/v1/posts", middleware.AuthMiddleware, controllers.PostsIndex)
	r.GET("/api/v1/posts/:postId", middleware.AuthMiddleware, controllers.PostsDetail)
	r.PUT("/api/v1/posts/:postId", middleware.AuthMiddleware, controllers.PostsUpdate)
	r.DELETE("/api/v1/posts/:postId", middleware.AuthMiddleware, controllers.PostsDelete)

	r.Run()
}
