package httpserver

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
)

func StartRouter() {

	config := Config{}
	if err := env.Parse(&config); err != nil {
		log.Fatalf("%+v", err)
	}

	router := gin.Default()

	// router.POST("/posts", service.DB.CreateUser)
	// router.PUT("/posts/:id", UpdateUser)
	// router.GET("/posts", GetAllUsers)
	// router.GET("/posts/:id", GetUserById)
	// router.DELETE("/posts/:id", DeleteUser)

	router.Run()
}
