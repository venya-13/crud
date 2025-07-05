package httpserver

import (
	"crud/internal/service"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	FamilyId  uint      `json:"family_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Family struct {
	FamilyId uint   `json:"family_id"`
	Name     string `json:"name"`
}

type FamiltMember struct {
	FamilyId uint `json:"family_id"`
	UserId   uint
	Role     string
	//Think what do you need int or uint
}

// add password field to User struct
// Needed functions Register, Login.

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
		Email:   post.Email,
		Age:     post.Age,
	})

	if err != nil {
		log.Println("CreateUser: handling database error in http:", err)
	}

	post.Id = id

	ginContext.JSON(http.StatusOK, gin.H{"post": post})

}

func (s *Server) CreateFamily(ginContext *gin.Context) {
	var family Family

	if err := ginContext.BindJSON(&family); err != nil {
		log.Println("Bind error", err)
	}

	if family.Name == "" {
		ginContext.JSON(http.StatusBadRequest, gin.H{"Error, family name field is empty ": family})
		return
	}

	// Check if family name already exists

	id, err := s.svc.CreateFamily(family.Name)
	if err != nil {
		log.Println("CreateFamily: handling database error in http:", err)
		ginContext.JSON(http.StatusInternalServerError, gin.H{"Failed to create family": "Http error"})
		return
	}

	family.FamilyId = id
	ginContext.JSON(http.StatusOK, gin.H{"Family: ": family})
}

func (s *Server) AddToFamily(ginContext *gin.Context) {
	var familyMember FamiltMember

	if err := ginContext.BindJSON(&familyMember); err != nil {
		log.Println("Bind error", err)
		ginContext.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to bind family member"})
		return
	}

	// Is Role field obligatory? If so, check it
	if familyMember.FamilyId == 0 || familyMember.UserId == 0 {
		ginContext.JSON(http.StatusBadRequest, gin.H{"Error": "FamilyId and UserId must be provided"})
		return
	}

	err := s.svc.AddToFamily(familyMember.FamilyId, familyMember.UserId, familyMember.Role)

	if err != nil {
		log.Println("AddToFamily: handling database error in http:", err)
		ginContext.JSON(http.StatusInternalServerError, gin.H{"Failed to add user to family": "Http error"})
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{"message": "User added to family successfully", "familyMember": familyMember})
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
			Email:   user.Email,
			Age:     user.Age,
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
			Email:   user.Email,
			Age:     user.Age,
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
		Email   string
		Age     int
	}

	if err := ginContext.BindJSON(&body); err != nil {
		log.Println("Bind error", err)
	}

	updatedUser, err := s.svc.UpdateUser(id, &service.User{
		Id:      body.Id,
		Name:    body.Name,
		Surname: body.Surname,
		Email:   body.Email,
		Age:     body.Age,
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
			Email:   user.Email,
			Age:     user.Age,
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
