package redisclient

import (
	"context"
	"crud/internal/service"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

type Redis struct {
	port    int
	context context.Context
	client  *redis.Client
}

type Service interface {
	GetUserById(id string) (*service.User, error)
	SaveUser(user *service.User) error
	DeleteUpdateUser(id string) error
}

func InitRedis() *Redis {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // default DB
	})

	r := Redis{
		port:    6379,
		context: ctx,
		client:  rdb,
	}

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	log.Printf("Connected to Redis: %s", pong)

	return &r
}
