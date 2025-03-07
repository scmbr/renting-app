package service

import (
	"github.com/vasya/renting-app/internal/dto"
	"github.com/vasya/renting-app/internal/models"
	"github.com/vasya/renting-app/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int,error)
	GenerateToken(email string,password string ) (string, error)
	ParseToken(token string) (int,error)
}
type Users interface {
	GetAllUsers() ([]dto.GetUser,error)
	GetUserById(id int) (*dto.GetUser,error)
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