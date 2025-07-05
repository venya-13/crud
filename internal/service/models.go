package service

import "time"

type User struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	FamilyId  uint      `json:"family_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Family struct {
	FamilyId uint   `json:"family_id"`
	Name     string `json:"name"`
}

type FamiltMember struct {
	FamilyId uint `json:"family_id"`
	UserId   uint
	Role     string
	//Think what do you need int or uint
}
