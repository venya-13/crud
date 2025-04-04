package main

import (
	httpserver "crud/internal/http-server"
	postgresdb "crud/internal/postgres-db"
	"fmt"
)

func main() {
	fmt.Println("Start")

	postgresdb.ConnectToDB()

	httpserver.StartRouter()
}
