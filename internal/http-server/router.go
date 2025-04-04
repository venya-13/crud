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

	router.POST("/posts", CreateUserHttp)
	router.PUT("/posts/:id", UpdateUserHttp)
	router.GET("/posts", GetAllUsersHttp)
	router.GET("/posts/:id", GetUserByIdHttp)
	router.DELETE("/posts/:id", DeleteUserHttp)

	router.Run(config.Port)
}
