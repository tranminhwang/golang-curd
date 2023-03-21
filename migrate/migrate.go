package main

import (
	"rest-api-with-golang/initializers"
	"rest-api-with-golang/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{}, &models.User{})
}
