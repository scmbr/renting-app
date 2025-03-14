package service

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/vasya/renting-app/internal/dto"
	"github.com/vasya/renting-app/internal/repository"
)

type Authorization interface {
	CreateUser(user dto.CreateUser) (int,error)
	GenerateToken(email string,password string ) (string, error)
	ParseToken(token string) (int,error)
}
type Users interface {
	GetAllUsers() ([]dto.GetUser,error)
	GetUserById(id int) (*dto.GetUser,error)
	GetCurrentUserId(с *gin.Context) (int,error)
	UploadAvatarToS3(fileHeader *multipart.FileHeader) (string,error)
	UpdateAvatar(userId int, avatarURL string)(error)
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