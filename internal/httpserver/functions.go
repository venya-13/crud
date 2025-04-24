package httpserver

import (
	"crud/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (s *Server) CreateUser(ginContext *gin.Context) {
	var post User

	if err := ginContext.BindJSON(&post); err != nil {
		log.Println("Bind error", err)
	}

	if post.Name == "" {
		ginContext.JSON(http.StatusBadRequest, gin.H{"Error": post})
	}

	id, err := s.svc.CreateUser(&service.User{
		Name:    post.Name,
		Surname: post.Surname,
	})

	if err != nil {
		log.Println("CreateUser: handling database error in http:", err)
	}

	post.Id = id

	ginContext.JSON(http.StatusOK, gin.H{"post": post})

}

func (s *Server) GetAllUsers(ginContext *gin.Context) {
	var posts []User
	users, err := s.svc.GetAllUsers()

	if err != nil {
		log.Println("GetAllUsers: handling database error in http:", err)
		ginContext.JSON(http.StatusInternalServerError, gin.H{"Failed to get user from database": "Http error"})
		return
	}

	for _, user := range users {
		posts = append(posts, User{
			Id:      user.Id,
			Name:    user.Name,
			Surname: user.Surname,
		})
	}

	ginContext.JSON(http.StatusOK, gin.H{"post": posts})

}

func (s *Server) GetUserById(ginContext *gin.Context) {

	var post []User

	id := ginContext.Param("id")

	users, err := s.svc.GetUserById(id)

	if err != nil {
		log.Println("GetUserById: handling database error in http:", err)
		ginContext.JSON(http.StatusInternalServerError, gin.H{"Failed to get user from database": "Http error"})
	}

	for _, user := range users {
		post = append(post, User{
			Id:      user.Id,
			Name:    user.Name,
			Surname: user.Surname,
		})
	}

	ginContext.JSON(http.StatusOK, gin.H{"post": post})

}

func (s *Server) UpdateUser(ginContext *gin.Context) {

	var post []User

	id := ginContext.Param("id")

	var body struct {
		Id      uint
		Name    string
		Surname string
	}

	if err := ginContext.BindJSON(&body); err != nil {
		log.Println("Bind error", err)
	}

	updatedUser, err := s.svc.UpdateUser(id, &service.User{
		Id:      body.Id,
		Name:    body.Name,
		Surname: body.Surname,
	})

	if err != nil {
		log.Println("UpdateUser: handling database error in http:", err)
		ginContext.JSON(http.StatusInternalServerError, gin.H{"Failed to update user from database": "Http error"})

	}

	for _, user := range updatedUser {
		post = append(post, User{
			Id:      user.Id,
			Name:    user.Name,
			Surname: user.Surname,
		})

		ginContext.JSON(http.StatusOK, gin.H{"post": post})
	}
}

func (s *Server) DeleteUser(ginContext *gin.Context) {
	id := ginContext.Param("id")

	err := s.svc.DeleteUser(id)

	if err != nil {
		log.Println("DeleteUser: handling database error in http:", err)
		ginContext.JSON(http.StatusInternalServerError, gin.H{"Failed to delete user from database": "Http error"})
		return
	}

	ginContext.Status(http.StatusOK)

}
