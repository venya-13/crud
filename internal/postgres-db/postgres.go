package postgresdb

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	DB struct {
		db *gorm.DB
	}
	Config struct {
		Port  string `envconfig:"HTTP_SERVER_PORT" default:"5431"`
		DBUrl string `envconfig:"DB_URL" default:"host=localhost user=postgres password=pass dbname=postgres port=5431 sslmode=disable"`
	}
)

func New(config Config) (*DB, error) {
	dsn := config.DBUrl

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database ", err)
	}

	return &DB{
		db: db,
	}, nil
}
