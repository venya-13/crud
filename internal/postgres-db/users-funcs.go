package postgresdb

import (
	"crud/internal/postgres-db/models"
	"crud/internal/service"
	"log"
)

func (db *DB) CreateUser(name string, surname string) error {

	log.Println("Creating user")

	post := models.User{Name: name, Surname: surname}

	result := db.db.Create(&post)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (db *DB) GetAllUsers() []service.User {

	var posts []service.User
	db.db.Find(&posts)

	return posts

}

func (db *DB) GetUserById(id string) []service.User {

	var user []service.User
	db.db.First(&user, id)

	return user

}

func (db *DB) UpdateUser(id string, user service.User) []service.User {

	var post []service.User
	db.db.First(&post, id)

	db.db.Model(&post).Updates(user)

	return post
}

func (db *DB) DeleteUser(id string) {

	db.db.Delete(&service.User{}, id)
}
