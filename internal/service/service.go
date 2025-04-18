package service

type Service struct {
	db DB
}

type DB interface {
	CreateUser(name string, surname string, id uint) error
	GetAllUsers() ([]User, error)
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

	return err
}

func (svc *Service) GetAllUsers() ([]User, error) {

	posts, err := svc.db.GetAllUsers()

	return posts, err
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
