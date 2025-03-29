package migrate

import (
	"crud/internal/initializers"
	"crud/internal/models"
)

func Migrate() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	initializers.DB.AutoMigrate(&models.User{})
}
