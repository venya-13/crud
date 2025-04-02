package cmd

import (
	httpserver "crud/internal/http-server"
	postgresdb "crud/internal/postgres-db"
	"fmt"
	//"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Start")

	postgresdb.ConnectToDB()

	postgresdb.Migrate()

	httpserver.StartRouter()
}
