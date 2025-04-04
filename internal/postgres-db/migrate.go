package postgresdb

import (
	"crud/internal/service"
)

func Migrate(user *service.User) {

	ConnectToDB()

	db.AutoMigrate(&user)
}
