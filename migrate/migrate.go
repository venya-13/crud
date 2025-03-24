package migrate

import (
	"crud/initializers"
	"crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func Migrate() {
	initializers.DB.AutoMigrate(&models.Post{})
}
