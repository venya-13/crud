package postgresdb

import (
	"crud/internal/postgres-db/models"
)

// Problems :
// gin-context.bind problem

func CreateUser(name string, surname string) (models.User, error) {

	post := models.User{Name: name, Surname: surname}

	result := db.Create(&post)

	if result.Error != nil {
		return post, result.Error
	}

	return post, nil

}

func GetAllUsers() []models.User {

	var posts []models.User
	db.Find(&posts)

	return posts

}

func GetUserById(id string) []models.User {

	var post []models.User
	db.First(&post, id)

	return post

}

func UpdateUser(id string, name string, surname string) []models.User {

	var post []models.User
	db.First(&post, id)

	db.Model(&post).Updates(models.User{
		Name:    name,
		Surname: surname,
	})

	return post
}

func DeleteUser(id string) {

	db.Delete(&models.User{}, id)
}
