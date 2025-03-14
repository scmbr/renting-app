package service

import (
	"errors"
	"mime/multipart"

	"github.com/gin-gonic/gin"
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
func (s *UsersService) GetCurrentUserId(c *gin.Context) (int,error){
	
	userID, exists := c.Get("userCtx")
    if !exists {
        return 0, errors.New("user ID not found in context")
    }

    id, ok := userID.(int)
    if !ok {
        return 0, errors.New("invalid user ID type")
    }

    return id, nil
}
func (s *UsersService) UploadAvatarToS3(fileHeader *multipart.FileHeader) (string,error){
	
	return "1",nil
}
func (s *UsersService) UpdateAvatar(userId int,avatarURL string) (error){
	
	return nil
}
