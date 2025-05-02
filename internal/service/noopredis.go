package service

type NoopRedis struct{}

func (n *NoopRedis) GetUserById(id string) (*User, error) {
	return nil, nil
}

func (n *NoopRedis) SaveUser(user *User) error {
	return nil
}

func (n *NoopRedis) DeleteUpdateUser(id string) error {
	return nil
}

func (n *NoopRedis) GetAllUsers() ([]User, error) {
	return nil, nil
}

func (n *NoopRedis) SaveAllUsers(users []User) error {
	return nil
}
