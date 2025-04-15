package httpserver

import (
	"crud/internal/http-server/models"
	"crud/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateUser(ginContext *gin.Context) {
	var post models.User

	if err := ginContext.BindJSON(&post); err != nil {
		fmt.Println("Bind error", err)
	}

	if post.Id != 0 && post.Name != "" {
		err := s.svc.CreateUser(&service.User{
			Id:      post.Id,
			Name:    post.Name,
			Surname: post.Surname,
		})

		if err != nil {
			fmt.Println("CreateUser: from http to server error:", err)
		}

		ginContext.JSON(200, gin.H{"post": post})
	} else {
		ginContext.JSON(400, gin.H{"Error": post})

	}

}

func (s *Server) GetAllUsers(ginContext *gin.Context) {
	var posts []models.User
	users, err := s.svc.GetAllUsers()

	if err != nil {
		fmt.Println("GetAllUsers: from http to server error:", err)
		ginContext.JSON(500, gin.H{"Failed to get user from database": "Http error"})
		return
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

	users, err := s.svc.GetUserById(id)

	if err != nil {
		fmt.Println("GetUserById: from http to server error:", err)
		ginContext.JSON(500, gin.H{"Failed to get user from database": "Http error"})
	}

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

	if err := ginContext.BindJSON(&body); err != nil {
		fmt.Println("Bind error", err)
	}

	updatedUser, err := s.svc.UpdateUser(id, &service.User{
		Id:      body.Id,
		Name:    body.Name,
		Surname: body.Surname,
	})

	if err != nil {
		fmt.Println("UpdateUser: from http to server error:", err)
		ginContext.JSON(500, gin.H{"Failed to update user from database": "Http error"})

	}

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

	err := s.svc.DeleteUser(id)

	if err != nil {
		fmt.Println("DeleteUser:Failed to delete user error:", err)
		ginContext.JSON(500, gin.H{"Failed to delete user from database": "Http error"})
		return
	}

	ginContext.Status(200)

}
