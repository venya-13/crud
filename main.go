package main

import (
	"crud/controllers"
	"crud/initializers"
	"crud/migrate"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Start")

	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	migrate.Migrate()

	router := gin.Default()

	router.POST("/posts", controllers.PostUser)
	router.PUT("/posts/:id", controllers.UpdateUser)
	router.GET("/posts", controllers.GetAllUsers)
	router.GET("/posts/:id", controllers.GetUserById)
	router.DELETE("/posts/:id", controllers.DeleteUser)

	router.Run()
}
