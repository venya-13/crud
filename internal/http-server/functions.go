package httpserver

import (
	"crud/internal/http-server/models"
	"crud/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateUser(ginContext *gin.Context) {
	var post models.User

	var body struct {
		Name    string
		Surname string
	}

	ginContext.Bind(&body)

	err := s.svc.CreateUser(&service.User{
		Name:    body.Name,
		Surname: body.Surname,
	})

	if err != nil {
		log.Println("CreateUser: database error")
	}

	user := models.User{
		Name:    body.Name,
		Surname: body.Surname,
	}

	post = models.User(user)
	ginContext.JSON(200, gin.H{"post": post})
}

func (s *Server) GetAllUsers(ginContext *gin.Context) {
	var posts []models.User
	users := s.svc.GetAllUsers()

	for _, user := range users {
		posts = append(posts, models.User{
			Name:    user.Name,
			Surname: user.Surname,
		})
	}

	ginContext.JSON(200, gin.H{"post": posts})

}

func (s *Server) GetUserById(ginContext *gin.Context) {

	var post []models.User

	id := ginContext.Param("id")

	users := s.svc.GetUserById(id)

	for _, user := range users {
		post = append(post, models.User{
			Name:    user.Name,
			Surname: user.Surname,
		})
	}

	ginContext.JSON(200, gin.H{"post": post})

}

func (s *Server) UpdateUser(ginContext *gin.Context) {

	var post []models.User

	id := ginContext.Param("id")

	var body struct {
		Name    string
		Surname string
	}

	ginContext.Bind(&body)

	updatedUser := s.svc.UpdateUser(id, body.Name, body.Surname)

	for _, user := range updatedUser {
		post = append(post, models.User{
			Name:    user.Name,
			Surname: user.Surname,
		})

		ginContext.JSON(200, gin.H{"post": post})
	}
}

func (s *Server) DeleteUser(ginContext *gin.Context) {
	id := ginContext.Param("id")

	service.DB.DeleteUser(nil, id)

	ginContext.Status(200)
}
