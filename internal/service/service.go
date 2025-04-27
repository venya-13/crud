package service

import (
	"fmt"
)

type Service struct {
	db    DB
	redis Redis
}

type DB interface {
	CreateUser(name string, surname string) (uint, error)
	GetAllUsers() ([]User, error)
	GetUserById(id string) ([]User, error)
	UpdateUser(id string, user User) ([]User, error)
	DeleteUser(id string) error
	Close()
}

type Redis interface {
	GetUserById(id string) (*User, error)
	SaveUser(user *User) error
	DeleteUpdateUser(id string) error
}

func New(db DB, redis Redis) *Service {

	if redis == nil {
		redis = &NoopRedis{}
	}

	return &Service{
		db:    db,
		redis: redis,
	}
}

func (svc *Service) CreateUser(user *User) (uint, error) {
	id, err := svc.db.CreateUser(user.Name, user.Surname)

	if err != nil {
		return 0, fmt.Errorf("create user database error %w", err)
	}
	return id, nil
}

func (svc *Service) GetAllUsers() ([]User, error) {

	posts, err := svc.db.GetAllUsers()

	if err != nil {
		return posts, fmt.Errorf("get all users error %w", err)
	}
	return posts, nil
}

func (svc *Service) GetUserById(id string) ([]User, error) {

	user, err := svc.redis.GetUserById(id)

	if err != nil {
		return nil, fmt.Errorf("get user by id from Redis error %w", err)
	}

	if user != nil {
		return []User{*user}, nil
	}

	userById, err := svc.db.GetUserById(id)
	if err != nil {
		return userById, fmt.Errorf("get user by id error %w", err)
	}

	if svc.redis != nil && len(userById) > 0 {
		err := svc.redis.SaveUser(&userById[0])
		if err != nil {
			return userById, fmt.Errorf("save user to Redis error %w", err)
		}
	}

	return userById, nil
}

func (svc *Service) UpdateUser(id string, user *User) ([]User, error) {

	updatedUser, err := svc.db.UpdateUser(id, *user)

	if err != nil {
		return updatedUser, fmt.Errorf("update user error %w", err)
	}

	if svc.redis != nil {
		err := svc.redis.DeleteUpdateUser(id)
		if err != nil {
			return updatedUser, fmt.Errorf("update user redis error %w", err)
		}
	}

	return updatedUser, nil
}

func (svc *Service) DeleteUser(id string) error {
	err := svc.db.DeleteUser(id)

	if err != nil {
		return fmt.Errorf("delete user error %w", err)
	}

	if svc.redis != nil {
		err := svc.redis.DeleteUpdateUser(id)

		if err != nil {
			return fmt.Errorf("delete user redis error %w", err)
		}
	}

	return nil
}
