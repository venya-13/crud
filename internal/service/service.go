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

	if errFinal := fmt.Errorf("create user database error %w", err); err != nil {
		return errFinal
	}
	return nil
}

func (svc *Service) GetAllUsers() ([]User, error) {

	posts, err := svc.db.GetAllUsers()

	if errFinal := fmt.Errorf("get all users error %w", err); err != nil {
		return posts, errFinal
	}
	return posts, nil
}

func (svc *Service) GetUserById(id string) ([]User, error) {
	userById, err := svc.db.GetUserById(id)

	if errFinal := fmt.Errorf("get user by id error %w", err); err != nil {
		return userById, errFinal
	}
	return userById, nil
}

func (svc *Service) UpdateUser(id string, user *User) ([]User, error) {

	updatedUser, err := svc.db.UpdateUser(id, *user)

	if errFinal := fmt.Errorf("update user error %w", err); err != nil {
		return updatedUser, errFinal
	}
	return updatedUser, nil

}

func (svc *Service) DeleteUser(id string) error {
	err := svc.db.DeleteUser(id)

	if errFinal := fmt.Errorf("delete user error %w", err); err != nil {
		return errFinal
	}
	return nil
}
