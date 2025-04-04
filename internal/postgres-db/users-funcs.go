package postgresdb

import (
	"crud/internal/service"

	"github.com/gin-gonic/gin"
)

func CreateUser(ginContext *gin.Context, user *service.User) {

	var body struct {
		Name    string
		Surname string
	}

	ginContext.Bind(&body)

	//post := service.User{Name: body.Name, Surname: body.Surname}
	*user = service.User{Name: body.Name, Surname: body.Surname}

	result := db.Create(&user)

	if result.Error != nil {
		ginContext.Status(500)
		return
	}

	ginContext.JSON(200, gin.H{"post": user})
}

func GetAllUsers(ginContext *gin.Context, user *service.User) {

	var posts []service.User
	db.Find(&posts)

	ginContext.JSON(200, gin.H{"post": posts})

}

func GetUserById(ginContext *gin.Context) {

	id := ginContext.Param("id")

	var post []service.User
	db.First(&post, id)

	ginContext.JSON(200, gin.H{"post": post})

}

func UpdateUser(ginContext *gin.Context) {
	id := ginContext.Param("id")

	var body struct {
		Name    string
		Surname string
	}

	ginContext.Bind(&body)

	var post []service.User
	db.First(&post, id)

	db.Model(&post).Updates(service.User{
		Name:    body.Name,
		Surname: body.Surname,
	})

	ginContext.JSON(200, gin.H{"post": post})
}

func DeleteUser(ginContext *gin.Context) {
	id := ginContext.Param("id")

	db.Delete(&service.User{}, id)

	ginContext.Status(200)
}
