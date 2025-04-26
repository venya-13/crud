package redisclient

import (
	"crud/internal/service"
)

func (r *Redis) GetUserById(id string) (*service.User, error) {
	// Implement the logic to get the user by ID from Redis
	return nil, nil

}

func (r *Redis) SaveUser(user *service.User) error {
	// Implement the logic to save the user in Redis
	return nil
}

func (r *Redis) DeleteUpdateUser(id string) error {
	// Implement the logic to delete or update the user in Redis
	return nil
}
