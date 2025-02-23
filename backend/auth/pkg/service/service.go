package service

import (
	"github.com/vasya/auth"
	"github.com/vasya/auth/pkg/repository"
)

type Authorization interface {
	CreateUser(user auth.User) (int,error)
}
type Services struct {
	Authorization
}

func NewServices(repos *repository.Repository) *Services {
	return &Services{
		Authorization: NewAuthService(repos.Authorization),
	}
}