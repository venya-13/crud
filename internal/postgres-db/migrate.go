package postgresdb

import "crud/internal/postgres-db/models"

func Migrate() {

	ConnectToDB()

	db.AutoMigrate(&models.User{})
}
