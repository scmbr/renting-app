package repository

import (
	"context"
	"time"

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
	CreateUser(ctx context.Context, user dto.CreateUser, code string) error
	GetUser(email, password string) (models.User, error)
	GetByCredentials(ctx context.Context, email, password string) (*dto.GetUser, error)
	Verify(ctx context.Context, code string) (dto.GetUser, error)
	GetByEmail(ctx context.Context, email string) (*dto.GetUser, error)
	UpdateVerificationCode(ctx context.Context, id int, verificationCode string) error
	SavePasswordResetToken(ctx context.Context, id int, resetToken string) error
	GetUserByResetToken(ctx context.Context, token string) (dto.GetUser, error)
	UpdatePasswordAndClearResetToken(ctx context.Context, userID int, newPassword string) error
}
type Session interface {
	CreateSession(ctx context.Context, session models.Session) error
	GetByRefreshToken(ctx context.Context, refreshToken string) (models.Session, error)
	UpdateSession(ctx context.Context, session models.Session) error
	GetByDevice(ctx context.Context, userID int, ip, os, browser string) (*models.Session, error)
	UpdateTokens(ctx context.Context, sessionID int, refreshToken string, expiresAt time.Time) error
	DeleteByDevice(ctx context.Context, id int, ip, os, browser string) error
}
type Apartment interface {
	GetAllApartments(ctx context.Context, userId int) ([]*dto.GetApartmentResponse, error)
	GetApartmentById(ctx context.Context, userId int, id int) (*dto.GetApartmentResponse, error)
	CreateApartment(ctx context.Context, userId int, input dto.CreateApartmentInput) error
	DeleteApartment(ctx context.Context, userId int, id int) error
	UpdateApartment(ctx context.Context, userId int, id int, input *dto.UpdateApartmentInput) error
}
type Advert interface {
	GetAllAdverts(ctx context.Context, userId int) ([]*dto.GetAdvertResponse, error)
	GetAdvertById(ctx context.Context, userId int, id int) (*dto.GetAdvertResponse, error)
	CreateAdvert(ctx context.Context, userId int, input dto.CreateAdvertInput) error
	DeleteAdvert(ctx context.Context, userId int, id int) error
	UpdateAdvert(ctx context.Context, userId int, id int, input *dto.UpdateAdvertInput) error
}
type Repository struct {
	Users
	Session
	Apartment
	Advert
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Users:     NewUsersRepo(db),
		Session:   NewSessionsRepo(db),
		Apartment: NewApartmentRepo(db),
		Advert:    NewAdvertRepo(db),
	}
}
