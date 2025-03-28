package migrate

import (
	"crud/initializers"
	"crud/models"
)

func Migrate() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	initializers.DB.AutoMigrate(&models.User{})
}
