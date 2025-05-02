package service

import "time"

type User struct {
	Id        uint      `note:"column: id"`
	Name      string    `note:"column: name"`
	Surname   string    `note:"column: surname"`
	UpdatedAt time.Time `note:"column: updated_at"`
}
