package controllers

import (
	"github.com/gin-gonic/gin"
	"rest-api-with-golang/initializers"
	"rest-api-with-golang/models"
)

func PostsCreate(c *gin.Context) {
	var requestBody struct {
		Title string `form:"title"`
		Body  string `form:"body"`
	}
	c.Bind(&requestBody)

	post := models.Post{
		Title: requestBody.Title,
		Body:  requestBody.Body,
	}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(201, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsDetail(c *gin.Context) {
	var post models.Post
	postId := c.Param("postId")

	initializers.DB.First(&post, postId)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	var requestBody struct {
		Title string `form:"title"`
		Body  string `form:"body"`
	}
	var post models.Post

	c.Bind(&requestBody)
	postId := c.Param("postId")
	// find post
	initializers.DB.First(&post, postId)
	// update
	initializers.DB.Model(&post).Updates(models.Post{
		Title: requestBody.Title,
		Body:  requestBody.Body,
	})

	c.JSON(200, gin.H{
		"message": "update success",
	})

}
func PostsDelete(c *gin.Context) {
	postId := c.Param("postId")

	initializers.DB.Delete(&models.Post{}, postId)

	c.JSON(200, gin.H{
		"message": "delete success",
	})
}
