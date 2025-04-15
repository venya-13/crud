package postgresdb

import (
	"crud/internal/postgres-db/models"
	"crud/internal/service"
	"fmt"
	"log"
)

func (db *DB) CreateUser(name string, surname string, id uint) error {

	log.Println("Creating user")

	post := models.User{Name: name, Surname: surname, Id: id}

	result := db.db.Create(&post)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (db *DB) GetAllUsers() ([]service.User, error) {

	var posts []service.User
	errorCheck := db.db.Find(&posts)

	fmt.Println(posts)

	if errorCheck.Error != nil {
		return posts, errorCheck.Error
	}

	return posts, nil

}

func (db *DB) GetUserById(id string) ([]service.User, error) {

	var post []service.User
	errorCheck := db.db.First(&post, id)

	if errorCheck.Error != nil {
		return post, errorCheck.Error
	}

	fmt.Println(post)

	return post, nil

}

func (db *DB) UpdateUser(id string, user service.User) ([]service.User, error) {

	var post []service.User
	db.db.First(&post, id)

	errorCheck := db.db.Model(&post).Updates(user)

	if errorCheck.Error != nil {
		return post, errorCheck.Error
	}

	return post, nil
}

func (db *DB) DeleteUser(id string) error {

	errorCheck := db.db.Delete(&service.User{}, id)

	if errorCheck.Error != nil {
		return errorCheck.Error
	}

	return nil
}
