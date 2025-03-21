package service

import (
	"mime/multipart"

	"github.com/vasya/renting-app/internal/dto"
	"github.com/vasya/renting-app/internal/repository"
)

type Authorization interface {
	CreateUser(user dto.CreateUser) (int, error)
	GenerateToken(email string, password string) (string, error)
	ParseToken(token string) (int, error)
}
type Users interface {
	GetAllUsers() ([]dto.GetUser, error)
	GetUserById(id int) (*dto.GetUser, error)
	DeleteUserById(id int) (*dto.GetUser, error)
	UpdateUserById(input *dto.GetUser) (*dto.GetUser, error)
	UploadAvatarToS3(fileHeader *multipart.FileHeader) (string, error)
	UpdateAvatar(userId int, avatarURL string) error
}
type Services struct {
	Authorization
	Users
}

func NewServices(repos *repository.Repository) *Services {
	return &Services{
		Authorization: NewAuthService(repos.Authorization),
		Users:         NewUsersService(repos.Users),
	}
}
