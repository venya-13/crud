package service

import "github.com/gin-gonic/gin"

// Make http interface

type Service struct {
	db DB
}

type DB interface {
	CreateUser(ginContext *gin.Context) (User, error)
	GetAllUsers(ginContext *gin.Context) []User
	GetUserById(ginContext *gin.Context) []User
	UpdateUser(ginContext *gin.Context) []User
}
