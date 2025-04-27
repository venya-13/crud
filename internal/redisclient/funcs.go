package redisclient

import (
	"crud/internal/service"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func (r *Redis) GetUserById(id string) (*service.User, error) {
	key := fmt.Sprintf("user:%s", id)
	val, err := r.client.Get(r.context, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("could not get user from Redis: %w", err)
	}

	var user service.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Redis) SaveUser(user *service.User) error {
	key := fmt.Sprintf("user:%d", user.Id)

	bytes, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("could not marshal user: %w", err)
	}

	err = r.client.Set(r.context, key, bytes, 10*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("could not save user to Redis: %w", err)
	}

	return nil
}

func (r *Redis) DeleteUpdateUser(id string) error {
	key := fmt.Sprintf("user:%s", id)

	err := r.client.Del(r.context, key).Err()
	if err != nil {
		return fmt.Errorf("could not delete user from Redis: %w", err)
	}

	return nil
}
