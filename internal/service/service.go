package service

// Something wrong here

type Service struct {
	db DB
}

type DB interface {
	CreateUser(name string, surname string) (User, error)
	GetAllUsers() []User
	GetUserById(id string) []User
	UpdateUser(id string, name string, surname string) []User
	DeleteUser(id string)
}
