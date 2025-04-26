package service

import "fmt"

type Service struct {
	db DB
	r  Redis
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

func New(db DB, r Redis) *Service {
	return &Service{
		db: db,
		r:  r,
	}
}

func (svc *Service) CreateUser(user *User) (uint, error) {
	id, err := svc.db.CreateUser(user.Name, user.Surname)

	if errFinal := fmt.Errorf("create user database error %w", err); err != nil {
		return 0, errFinal
	}
	return id, nil
}

func (svc *Service) GetAllUsers() ([]User, error) {

	posts, err := svc.db.GetAllUsers()

	if errFinal := fmt.Errorf("get all users error %w", err); err != nil {
		return posts, errFinal
	}
	return posts, nil
}

func (svc *Service) GetUserById(id string) ([]User, error) {
	if svc.r != nil {

		user, err := svc.r.GetUserById(id)

		if err == nil && user != nil {
			return []User{*user}, nil
		} else {
			return nil, fmt.Errorf("get user by id error %w", err)
		}
	}

	userById, err := svc.db.GetUserById(id)
	if errFinal := fmt.Errorf("get user by id error %w", err); err != nil {
		return userById, errFinal
	}

	if svc.r != nil && len(userById) > 0 {
		_ = svc.r.SaveUser(&userById[0])
	}

	return userById, nil
}

func (svc *Service) UpdateUser(id string, user *User) ([]User, error) {

	updatedUser, err := svc.db.UpdateUser(id, *user)

	if errFinal := fmt.Errorf("update user error %w", err); err != nil {
		return updatedUser, errFinal
	}

	if svc.r != nil {
		_ = svc.r.DeleteUpdateUser(id)
	}

	return updatedUser, nil
}

func (svc *Service) DeleteUser(id string) error {
	err := svc.db.DeleteUser(id)

	if errFinal := fmt.Errorf("delete user error %w", err); err != nil {
		return errFinal
	}

	if svc.r != nil {
		_ = svc.r.DeleteUpdateUser(id)
	}

	return nil
}
