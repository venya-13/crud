package main

import (
	httpserver "crud/internal/httpserver"
	postgresdb "crud/internal/postgres-db"
)

type Config struct {
	HttpServer httpserver.Config `envconfig:"HTTP_SERVER_"`
	Postgres   postgresdb.Config `envconfig:"POSTGRES_"`
}
