package service

import (
	"fmt"
)

type Service struct {
	db    DB
	redis Redis
}

type DB interface {
	CreateUser(name, surname, email string, age int) (uint, error)
	GetAllUsers() ([]User, error)
	GetUserById(id string) ([]User, error)
	UpdateUser(id string, user User) ([]User, error)
	DeleteUser(id string) error
	Close()
	CreateFamily(familyName string) (uint, error)
	AddToFamily(familyId, userId uint, role string) error
}

type Redis interface {
	GetUserById(id string) (*User, error)
	SaveUser(user *User) error
	DeleteUpdateUser(id string) error
	GetAllUsers() ([]User, error)
	SaveAllUsers([]User) error
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
	id, err := svc.db.CreateUser(user.Name, user.Surname, user.Email, user.Age)

	if err != nil {
		return 0, fmt.Errorf("create user database error %w", err)
	}
	return id, nil
}

func (svc *Service) CreateFamily(familyName string) (uint, error) {
	id, err := svc.db.CreateFamily(familyName)
	if err != nil {
		return 0, fmt.Errorf("create family database error %w", err)
	}

	return id, nil
}

func (svc *Service) AddToFamily(familyId, userId uint, role string) error {
	err := svc.db.AddToFamily(familyId, userId, role)
	if err != nil {
		return fmt.Errorf("add to family database error %w", err)
	}

	return nil
}

func (svc *Service) GetAllUsers() ([]User, error) {

	cachedUsers, err := svc.redis.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("get users from redis error %w", err)
	}

	dbUsers, err := svc.db.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("get users from db error %w", err)
	}

	// if redis is empty, use db
	if len(cachedUsers) == 0 {
		_ = svc.redis.SaveAllUsers(dbUsers)
		return dbUsers, nil
	}

	dbIsNewer := false

	cacheMap := make(map[uint]User)

	for _, cachedUser := range cachedUsers {
		cacheMap[cachedUser.Id] = cachedUser
	}

	for _, dbUser := range dbUsers {
		if cachedUsers, ok := cacheMap[dbUser.Id]; !ok || dbUser.UpdatedAt.After(cachedUsers.UpdatedAt) {
			dbIsNewer = true
			break
		}
	}

	if dbIsNewer {
		_ = svc.redis.SaveAllUsers(dbUsers)
		return dbUsers, nil

	}

	return cachedUsers, nil
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
