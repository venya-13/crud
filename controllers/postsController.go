package controllers

import (
	"crud/initializers"
	"crud/models"

	"github.com/gin-gonic/gin"
)

func PostUser(ginContext *gin.Context) {

	var body struct {
		Name    string
		Surname string
	}

	ginContext.Bind(&body)

	//Create a post

	post := models.User{Name: body.Name, Surname: body.Surname}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ginContext.Status(500)
		return
	}

	ginContext.JSON(200, gin.H{"post": post})
}

func GetAllUsers(ginContext *gin.Context) {

	var posts []models.User
	initializers.DB.Find(&posts)

	ginContext.JSON(200, gin.H{"post": posts})

}

func GetUserById(ginContext *gin.Context) {

	id := ginContext.Param("id")

	var post []models.User
	initializers.DB.First(&post, id)

	ginContext.JSON(200, gin.H{"post": post})

}

func UpdateUser(ginContext *gin.Context) {
	id := ginContext.Param("id")

	var body struct {
		Name    string
		Surname string
	}

	ginContext.Bind(&body)

	var post []models.User
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.User{
		Name:    body.Name,
		Surname: body.Surname,
	})

	ginContext.JSON(200, gin.H{"post": post})
}

func DeleteUser(ginContext *gin.Context) {
	id := ginContext.Param("id")

	initializers.DB.Delete(&models.User{}, id)

	ginContext.Status(200)
}
