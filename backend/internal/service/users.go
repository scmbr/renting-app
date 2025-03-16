package service

import (
	"mime/multipart"

	"github.com/vasya/renting-app/internal/cloud"
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

func (s *UsersService) UploadAvatarToS3(fileHeader *multipart.FileHeader) (string,error){
	
	return cloud.UploadAvatarToS3(fileHeader)
}
func (s *UsersService) UpdateAvatar(userId int,avatarURL string) (error){
	
	return s.repo.UpdateAvatar(userId, avatarURL)
}
