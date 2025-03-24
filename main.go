package main

import (
	"crud/controllers"
	"crud/initializers"
	"crud/migrate"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("Hello")

	migrate.Migrate()

	router := gin.Default()

	router.POST("/posts", controllers.PostsCreate)
	router.PUT("/posts/:id", controllers.PostUpdate)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsShow)
	router.DELETE("/posts/:id", controllers.PostsDelete)

	router.Run()
}
