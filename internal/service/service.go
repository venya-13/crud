package service

import "github.com/gin-gonic/gin"

type Service struct {
	db DB
}

type DB interface {
	CreateUser(ginContext *gin.Context, u *User) error
	Migrate(user *User)
}
