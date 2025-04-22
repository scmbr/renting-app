package service

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/repository"
	"github.com/scmbr/renting-app/pkg/auth"
	"github.com/scmbr/renting-app/pkg/hash"
	"github.com/scmbr/renting-app/pkg/storage"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
type User interface {
	GetAllUsers() ([]dto.GetUser, error)
	GetUserById(id int) (*dto.GetUser, error)
	DeleteUserById(id int) (*dto.GetUser, error)
	UpdateUserById(input *dto.GetUser) (*dto.GetUser, error)
	UploadAvatarToS3(ctx context.Context, fileHeader *multipart.FileHeader) (string, error)
	UpdateAvatar(userId int, avatarURL string) error
	SignIn(ctx context.Context, email string, password string, ip string, os string, browser string) (Tokens, error)
	SignUp(ctx context.Context, user dto.CreateUser) error
}
type Session interface {
	CreateSession(ctx context.Context, userID int, ip string, os string, browser string) (Tokens, error)
	RefreshSession(ctx context.Context, refreshToken, ip, os, browser string) (Tokens, error)
}
type Services struct {
	User
	Session
}

type Deps struct {
	Repos           *repository.Repository
	Hasher          hash.PasswordHasher
	StorageProvider storage.Provider
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	TokenManager    auth.TokenManager
}

func NewServices(deps Deps) *Services {
	sessionService := NewSessionService(deps.Repos.Session, deps.AccessTokenTTL, deps.RefreshTokenTTL, deps.TokenManager)
	userService := NewUserService(
		deps.Repos.Users,
		deps.StorageProvider,
		deps.Hasher,
		deps.AccessTokenTTL,
		deps.RefreshTokenTTL,
		deps.TokenManager,
		sessionService)

	return &Services{
		User:    userService,
		Session: sessionService,
	}
}
