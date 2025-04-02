package httpserver

import "github.com/gin-gonic/gin"

func StartRouter() {
	router := gin.Default()

	// router.POST("/posts", PostUser)
	// router.PUT("/posts/:id", UpdateUser)
	// router.GET("/posts", GetAllUsers)
	// router.GET("/posts/:id", GetUserById)
	// router.DELETE("/posts/:id", DeleteUser)

	router.Run()
}
