package controllers

import (
	"crud/initializers"
	"crud/models"

	"github.com/gin-gonic/gin"
)

func PostUser(context *gin.Context) {

	var body struct {
		Name    string
		Surname string
	}

	context.Bind(&body)

	//Create a post

	post := models.User{Name: body.Name, Surname: body.Surname}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		context.Status(500)
		return
	}

	context.JSON(200, gin.H{"post": post})
}

func GetAllUsers(context *gin.Context) {

	var posts []models.User
	initializers.DB.Find(&posts)

	context.JSON(200, gin.H{"post": posts})

}

func GetUserById(context *gin.Context) {

	id := context.Param("id")

	var post []models.User
	initializers.DB.First(&post, id)

	context.JSON(200, gin.H{"post": post})

}

func UpdateUser(context *gin.Context) {
	id := context.Param("id")

	var body struct {
		Name    string
		Surname string
	}

	context.Bind(&body)

	var post []models.User
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.User{
		Name:    body.Name,
		Surname: body.Surname,
	})

	context.JSON(200, gin.H{"post": post})
}

func DeleteUser(context *gin.Context) {
	id := context.Param("id")

	initializers.DB.Delete(&models.User{}, id)

	context.Status(200)
}
