package service

import (
	"github.com/vasya/renting-app/internal/dto"
	"github.com/vasya/renting-app/internal/repository"
)


type UsersService struct {
	repo repository.Users
}


func NewUsersService(repo repository.Users) *UsersService{
	return &UsersService{repo:repo}
}

func (s *UsersService) GetAllUsers() ([]dto.GetUser,error){
	
	return s.repo.GetAllUsers()
}
func (s *UsersService) GetUserById(id int) (*dto.GetUser,error){
	
	return s.repo.GetUserById(id)
}
