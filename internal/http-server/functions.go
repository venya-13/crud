package httpserver

import (
	"crud/internal/http-server/models"
	"crud/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

var body struct {
	Name    string
	Surname string
}

func CreateUserHttp(ginContext *gin.Context, post models.User) {
	user, err := service.DB.CreateUser(nil, ginContext)

	if err != nil {
		log.Println("CreateUserHttp: getting from database error")
	}

	post = models.User(user)
	ginContext.JSON(200, gin.H{"post": post})
}

func GetAllUsersHttp(ginContext *gin.Context, posts []models.User) {
	users := service.DB.GetAllUsers(nil, ginContext)

	for _, user := range users {
		posts = append(posts, models.User{
			Name:    user.Name,
			Surname: user.Surname,
		})
	}

	ginContext.JSON(200, gin.H{"post": posts})

}

func GetUserByIdHttp(ginContext *gin.Context, post []models.User) {

	users := service.DB.GetAllUsers(nil, ginContext)

	for _, user := range users {
		post = append(post, models.User{
			Name:    user.Name,
			Surname: user.Surname,
		})
	}

	ginContext.JSON(200, gin.H{"post": post})

}

func UpdateUserHttp(ginContext *gin.Context, post []models.User) {

	users := service.DB.GetAllUsers(nil, ginContext)

	for _, user := range users {
		post = append(post, models.User{
			Name:    user.Name,
			Surname: user.Surname,
		})

		ginContext.JSON(200, gin.H{"post": post})
	}
}

func DeleteUserHttp(ginContext *gin.Context) {
	ginContext.Status(200)
}

// func Bind(ginContext *gin.Context) *body {

// 	ginContext.Bind(&body)

// 	return &body
// }
