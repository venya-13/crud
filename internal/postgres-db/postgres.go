package postgresdb

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	DB struct {
		db *pgxpool.Pool
	}
	Config struct {
		Port  string `envconfig:"HTTP_SERVER_PORT" default:"3000"`
		DBUrl string `envconfig:"DB_URL" default:"host=db user=postgres password=pass dbname=postgres port=5432 sslmode=disable"`
	}
)

func New(config Config) (*DB, error) {
	dbUrl := config.DBUrl

	db, err := pgxpool.Connect(context.Background(), dbUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &DB{
		db: db,
	}, nil
}

func dbClose(db *pgxpool.Pool) {
	db.Close()
}
