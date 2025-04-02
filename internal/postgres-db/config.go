package postgresdb

type DB struct {
	connectionString string
}

type Config struct {
	Port  string `env:"PORT" envDefault:"3000"`
	DBUrl string `env:"DB_URL" envDefault:"host=localhost user=postgres password=pass dbname=postgres port=5431 sslmode=disable"`
}
