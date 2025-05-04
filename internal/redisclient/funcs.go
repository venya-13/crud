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

	tx := r.client.TxPipeline()

	set := tx.Set(r.context, key, bytes, 10*time.Minute)
	_, err = tx.Exec(r.context)

	if err != nil {
		return fmt.Errorf("could not save user to redis(transaction failed): %w", err)
	}

	if set.Err() != nil {
		return fmt.Errorf("set command fails redis: %w", set.Err())
	}

	return nil
}

func (r *Redis) DeleteUpdateUser(id string) error {
	key := fmt.Sprintf("user:%s", id)

	tx := r.client.TxPipeline()

	del := tx.Del(r.context, key)
	_, err := tx.Exec(r.context)

	if err != nil {
		return fmt.Errorf("could not delete user from redis(transaction failed): %w", err)
	}

	if del.Err() != nil {
		return fmt.Errorf("delete command fails redis: %w", del.Err())
	}

	return nil
}

func (r *Redis) GetAllUsers() ([]service.User, error) {
	val, err := r.client.Get(r.context, "users:all").Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get all users from redis failed: %w", err)
	}

	var users []service.User
	if err := json.Unmarshal([]byte(val), &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Redis) SaveAllUsers(users []service.User) error {
	bytes, err := json.Marshal(users)
	if err != nil {
		return fmt.Errorf("marshal users for redis failed: %w", err)
	}

	err = r.client.Set(r.context, "users:all", bytes, 10*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("set all users to redis failed: %w", err)
	}
	return nil
}
