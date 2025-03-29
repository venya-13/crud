package cmd

import (
	"crud/internal/controllers"
	"crud/internal/initializers"
	"crud/migrate"
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
	//"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Start")

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v", err)
	}

	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	migrate.Migrate()

	// router := gin.Default()

	// router.POST("/posts", controllers.PostUser)
	// router.PUT("/posts/:id", controllers.UpdateUser)
	// router.GET("/posts", controllers.GetAllUsers)
	// router.GET("/posts/:id", controllers.GetUserById)
	// router.DELETE("/posts/:id", controllers.DeleteUser)

	// router.Run()

	controllers.StarListening()
}
