package models

type User struct {
	Id      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
