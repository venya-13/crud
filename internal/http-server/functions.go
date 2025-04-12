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
		Id      uint
		Name    string
		Surname string
	}

	ginContext.Bind(&body)

	if body.Id != 0 && body.Name != "" {
		err := s.svc.CreateUser(&service.User{
			Id:      body.Id,
			Name:    body.Name,
			Surname: body.Surname,
		})

		if err != nil {
			log.Fatal("CreateUser: from http to server error:", err)
		}

		user := models.User{
			Id:      body.Id,
			Name:    body.Name,
			Surname: body.Surname,
		}

		post = models.User(user)
		ginContext.JSON(200, gin.H{"post": post})
	} else {
		log.Println("User body is empty")
	}

}

func (s *Server) GetAllUsers(ginContext *gin.Context) {
	var posts []models.User
	users, err := s.svc.GetAllUsers()

	if err != nil {
		log.Fatal("GetAllUsers: from http to server error:", err)
	}

	for _, user := range users {
		posts = append(posts, models.User{
			Id:      user.Id,
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
			Id:      user.Id,
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
		Id      uint
		Name    string
		Surname string
	}

	ginContext.Bind(&body)

	updatedUser := s.svc.UpdateUser(id, &service.User{
		Id:      body.Id,
		Name:    body.Name,
		Surname: body.Surname,
	})

	for _, user := range updatedUser {
		post = append(post, models.User{
			Id:      user.Id,
			Name:    user.Name,
			Surname: user.Surname,
		})

		ginContext.JSON(200, gin.H{"post": post})
	}
}

func (s *Server) DeleteUser(ginContext *gin.Context) {
	id := ginContext.Param("id")

	s.svc.DeleteUser(id)

	ginContext.Status(200)
}
