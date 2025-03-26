package service

import (
	"mime/multipart"

	"github.com/vasya/renting-app/internal/dto"
	"github.com/vasya/renting-app/internal/repository"
	"github.com/vasya/renting-app/pkg/hash"
	"github.com/vasya/renting-app/pkg/storage"
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

type Deps struct {
	Repos           *repository.Repository
	Hasher          hash.PasswordHasher
	StorageProvider storage.Provider
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func NewServices(deps Deps) *Services {
	return &Services{
		Authorization: NewAuthService(deps.Repos.Authorization),
		Users:         NewUsersService(deps.Repos.Users),
	}
}
