package controllers

import (
	"crud/initializers"
	"crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(context *gin.Context) {

	var body struct {
		Name    string
		Surname string
	}

	context.Bind(&body)

	//Create a post

	post := models.Post{Name: body.Name, Surname: body.Surname}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		context.Status(400)
		return
	}

	context.JSON(200, gin.H{"post": post})
}

func PostsIndex(context *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)

	context.JSON(200, gin.H{"post": posts})

}

func PostsShow(context *gin.Context) {

	id := context.Param("id")

	var post []models.Post
	initializers.DB.First(&post, id)

	context.JSON(200, gin.H{"post": post})

}

func PostUpdate(context *gin.Context) {
	id := context.Param("id")

	var body struct {
		Name    string
		Surname string
	}

	context.Bind(&body)

	var post []models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Name:    body.Name,
		Surname: body.Surname,
	})

	context.JSON(200, gin.H{"post": post})
}

func PostsDelete(context *gin.Context) {
	id := context.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	context.Status(200)
}
