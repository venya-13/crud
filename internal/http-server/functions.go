package httpserver

import (
	"crud/internal/http-server/models"
	"crud/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

// Check, is it okay to send nil in serive.

func CreateUserHttp(ginContext *gin.Context) {
	var post models.User

	var body struct {
		Name    string
		Surname string
	}

	ginContext.Bind(&body)

	user, err := service.DB.CreateUser(nil, body.Name, body.Surname)

	if err != nil {
		log.Println("CreateUserHttp: getting from database error")
	}

	post = models.User(user)
	ginContext.JSON(200, gin.H{"post": post})
}

func GetAllUsersHttp(ginContext *gin.Context) {
	var posts []models.User
	users := service.DB.GetAllUsers(nil)

	for _, user := range users {
		posts = append(posts, models.User{
			Name:    user.Name,
			Surname: user.Surname,
		})
	}

	ginContext.JSON(200, gin.H{"post": posts})

}

func GetUserByIdHttp(ginContext *gin.Context) {

	var post []models.User

	id := ginContext.Param("id")

	users := service.DB.GetUserById(nil, id)

	for _, user := range users {
		post = append(post, models.User{
			Name:    user.Name,
			Surname: user.Surname,
		})
	}

	ginContext.JSON(200, gin.H{"post": post})

}

func UpdateUserHttp(ginContext *gin.Context) {

	var post []models.User

	id := ginContext.Param("id")

	var body struct {
		Name    string
		Surname string
	}

	ginContext.Bind(&body)

	users := service.DB.UpdateUser(nil, id, body.Name, body.Surname)

	for _, user := range users {
		post = append(post, models.User{
			Name:    user.Name,
			Surname: user.Surname,
		})

		ginContext.JSON(200, gin.H{"post": post})
	}
}

func DeleteUserHttp(ginContext *gin.Context) {
	id := ginContext.Param("id")

	service.DB.DeleteUser(nil, id)

	ginContext.Status(200)
}
