package service

type User struct {
	Id      uint   `note:"column: id"`
	Name    string `note:"column: name"`
	Surname string `note:"column: surname"`
}
