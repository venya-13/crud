package service

import "time"

type User struct {
	Id        uint      `note:"column: id"`
	Name      string    `note:"column: name"`
	Surname   string    `note:"column: surname"`
	Email     string    `note:"column: email"`
	Age       int       `note:"column: age"`
	UpdatedAt time.Time `note:"column: updated_at"`
}
