package postgresdb

import (
	"log"

	"github.com/caarlos0/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectToDB() {
	var err error

	config := Config{}
	if err := env.Parse(&config); err != nil {
		log.Fatalf("%+v", err)
	}

	dsn := config.DBUrl

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
