package postgresdb

import (
	"crud/internal/postgres-db/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(ginContext *gin.Context) (models.User, error) {

	var body struct {
		Name    string
		Surname string
	}

	ginContext.Bind(&body)

	post := models.User{Name: body.Name, Surname: body.Surname}

	result := db.Create(&post)

	if result.Error != nil {
		ginContext.Status(500)
		return post, result.Error
	}

	return post, nil

	// ginContext.JSON(200, gin.H{"post": post})
}

func GetAllUsers(ginContext *gin.Context) []models.User {

	var posts []models.User
	db.Find(&posts)

	return posts

	//ginContext.JSON(200, gin.H{"post": posts})

}

func GetUserById(ginContext *gin.Context) []models.User {

	id := ginContext.Param("id")

	var post []models.User
	db.First(&post, id)

	return post

	// ginContext.JSON(200, gin.H{"post": post})

}

func UpdateUser(ginContext *gin.Context) []models.User {
	id := ginContext.Param("id")

	var body struct {
		Name    string
		Surname string
	}

	ginContext.Bind(&body)

	var post []models.User
	db.First(&post, id)

	db.Model(&post).Updates(models.User{
		Name:    body.Name,
		Surname: body.Surname,
	})

	return post

	//ginContext.JSON(200, gin.H{"post": post})
}

func DeleteUser(ginContext *gin.Context) {
	id := ginContext.Param("id")

	db.Delete(&models.User{}, id)

	ginContext.Status(200)
}
