package repository

import (
	"context"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/models"
	"gorm.io/gorm"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
type Users interface {
	GetAllUsers() ([]dto.GetUser, error)
	GetUserById(id int) (*dto.GetUser, error)
	DeleteUserById(id int) (*dto.GetUser, error)
	UpdateUserById(input *dto.GetUser) (*dto.GetUser, error)
	UpdateAvatar(userId int, avatarURL string) error
	CreateUser(user dto.CreateUser) error
	GetUser(email, password string) (models.User, error)
	GetByCredentials(ctx context.Context, email, password string) (*dto.GetUser, error)
}
type Session interface {
	CreateSession(ctx context.Context, session models.Session) error
}
type Repository struct {
	Users
	Session
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Users:   NewUsersRepo(db),
		Session: NewSessionsRepo(db),
	}
}
