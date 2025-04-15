package service

import "fmt"

type Service struct {
	db DB
}

type DB interface {
	CreateUser(name string, surname string, id uint) error
	GetAllUsers() ([]User, error)
	GetUserById(id string) ([]User, error)
	UpdateUser(id string, user User) ([]User, error)
	DeleteUser(id string) error
}

func New(db DB) *Service {
	return &Service{
		db: db,
	}
}

func (svc *Service) CreateUser(user *User) error {
	err := svc.db.CreateUser(user.Name, user.Surname, user.Id)

	errFinal := fmt.Errorf("create user error %w", err)

	return errFinal
}

func (svc *Service) GetAllUsers() ([]User, error) {

	posts, err := svc.db.GetAllUsers()

	errFinal := fmt.Errorf("create user error %w", err)

	return posts, errFinal
}

func (svc *Service) GetUserById(id string) ([]User, error) {
	userById, err := svc.db.GetUserById(id)

	errFinal := fmt.Errorf("create user error %w", err)

	return userById, errFinal
}

func (svc *Service) UpdateUser(id string, user *User) ([]User, error) {

	updatedUser, err := svc.db.UpdateUser(id, *user)

	errFinal := fmt.Errorf("create user error %w", err)

	return updatedUser, errFinal
}

func (svc *Service) DeleteUser(id string) error {
	err := svc.db.DeleteUser(id)
	errFinal := fmt.Errorf("create user error %w", err)
	return errFinal
}
