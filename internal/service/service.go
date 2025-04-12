package service

import "fmt"

type Service struct {
	db DB
}

type DB interface {
	CreateUser(name string, surname string, id uint) error
	GetAllUsers() []User
	GetUserById(id string) []User
	UpdateUser(id string, user User) []User
	DeleteUser(id string)
}

func New(db DB) *Service {
	return &Service{
		db: db,
	}
}

func (svc *Service) CreateUser(user *User) error {
	err := svc.db.CreateUser(user.Name, user.Surname, user.Id)

	if err != nil {
		fmt.Println("Create user error", err)
	}

	return nil
}

func (svc *Service) GetAllUsers() []User {

	posts := svc.db.GetAllUsers()

	return posts
}

func (svc *Service) GetUserById(id string) []User {
	userById := svc.db.GetUserById(id)

	return userById
}

func (svc *Service) UpdateUser(id string, user *User) []User {

	updatedUser := svc.db.UpdateUser(id, *user)

	return updatedUser
}

func (svc *Service) DeleteUser(id string) {
	svc.db.DeleteUser(id)
}
