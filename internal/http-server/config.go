package httpserver

type Config struct {
	Port string `env:"PORT" envDefault:":3000"`
}
