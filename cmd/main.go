package cmd

import (
	httpserver "crud/internal/http-server"
	postgresdb "crud/internal/postgres-db"
	"crud/internal/service"
	"fmt"
	//"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Start")

	var db service.DB
	var u service.User

	postgresdb.ConnectToDB()

	service.DB.Migrate(db, &u)

	httpserver.StartRouter()
}
