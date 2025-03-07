package service

import (
	renting_app "github.com/vasya/renting-app"
	"github.com/vasya/renting-app/internal/repository"
)


type UsersService struct {
	repo repository.Users
}


func NewUsersService(repo repository.Users) *UsersService{
	return &UsersService{repo:repo}
}

func (s *UsersService) GetAllUsers() ([]renting_app.GetUser,error){
	
	return s.repo.GetAllUsers()
}
func (s *UsersService) GetUserById(id int) (*renting_app.GetUser,error){
	
	return s.repo.GetUserById(id)
}
