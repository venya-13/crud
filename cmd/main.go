package main

import (
	"crud/internal/httpserver"
	"crud/internal/postgresdb"
	"crud/internal/redisclient"
	"crud/internal/service"
	"fmt"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	fmt.Println("Start")

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	redis := redisclient.InitRedis()

	db, err := postgresdb.New(cfg.Postgres)

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}

	defer db.Close()

	svc := service.New(db, redis)

	httpServer := httpserver.New(cfg.HttpServer, svc)

	if err := httpServer.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
