package repository

import (
	renting_app "github.com/vasya/renting-app"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser (user renting_app.User) (int,error)
	GetUser(email,password string) (renting_app.User,error)
}
type Users interface{
	GetAllUsers() ([]renting_app.GetUser,error)
	GetUserById(id int) (*renting_app.GetUser,error)
}
type Repository struct {
	Authorization
	Users
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Users: NewUsersPostgres(db),
	}
}