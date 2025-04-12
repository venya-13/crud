package httpserver

import (
	"crud/internal/service"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Server struct {
		port   int
		router *gin.Engine
		svc    Service
	}

	Config struct {
		Port int `envconfig:"PORT" default:"3000"`
	}
)

type Service interface {
	CreateUser(user *service.User) error
	GetAllUsers() ([]service.User, error)
	GetUserById(id string) []service.User
	UpdateUser(id string, user *service.User) []service.User
	DeleteUser(id string)
}

func New(cfg Config, svc Service) *Server {
	s := Server{
		port:   cfg.Port,
		router: gin.Default(),
		svc:    svc,
	}

	s.router.POST("/posts", s.CreateUser)
	s.router.PUT("/posts/:id", s.UpdateUser)
	s.router.GET("/posts", s.GetAllUsers)
	s.router.GET("/posts/:id", s.GetUserById)
	s.router.DELETE("/posts/:id", s.DeleteUser)

	return &s
}

func (s *Server) Run() error {
	if err := s.router.Run(fmt.Sprintf(":%d", s.port)); !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to run server: %w", err)
	}

	return nil
}
