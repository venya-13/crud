package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockDB struct{}

func (m *mockDB) GetUserById(id string) ([]User, error) {
	return []User{
		{Id: 1, Name: "John", Surname: "Doe"},
	}, nil
}

func (m *mockDB) CreateUser(name, surname, email string, age int) (uint, error) { return 0, nil }
func (m *mockDB) CreateFamily(familyName string) (uint, error)                  { return 0, nil }
func (m *mockDB) AddToFamily(familyId, userId uint, role string) error          { return nil }
func (m *mockDB) GetAllUsers() ([]User, error)                                  { return nil, nil }
func (m *mockDB) UpdateUser(id string, user User) ([]User, error) {
	return nil, nil
}
func (m *mockDB) DeleteUser(id string) error { return nil }
func (m *mockDB) Close()                     {}

func TestGetUserById(t *testing.T) {
	db := &mockDB{}
	svc := New(db, nil)

	user, err := svc.GetUserById("1")

	assert.NoError(t, err)
	assert.Equal(t, 1, len(user)) // check is there one user or more
	assert.Equal(t, "John", user[0].Name)
	assert.Equal(t, "Doe", user[0].Surname)
}
