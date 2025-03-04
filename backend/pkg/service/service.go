package service

import (
	renting_app "github.com/vasya/renting-app"
	"github.com/vasya/renting-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user renting_app.User) (int,error)
	GenerateToken(email string,password string ) (string, error)
	ParseToken(token string) (int,error)
}
type Users interface {
	GetAllUsers() ([]renting_app.GetUser,error)
	GetUserById(id int) (*renting_app.GetUser,error)
}
type Services struct {
	Authorization
	Users
}

func NewServices(repos *repository.Repository) *Services {
	return &Services{
		Authorization: NewAuthService(repos.Authorization),
		Users: NewUsersService(repos.Users),
	}
}