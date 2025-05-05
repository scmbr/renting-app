package service

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/scmbr/renting-app/internal/config"
	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/repository"
	"github.com/scmbr/renting-app/pkg/auth"
	"github.com/scmbr/renting-app/pkg/email"
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
	VerifyEmail(ctx context.Context, code string) error
	ResendVerificationCode(ctx context.Context, email string) error
	ForgotPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, resetToken string, newPassword string) error
	LogOut(ctx context.Context, id int, ip, os, browser string) error
}
type Session interface {
	CreateSession(ctx context.Context, userID int, ip string, os string, browser string) (Tokens, error)
	RefreshSession(ctx context.Context, refreshToken, ip, os, browser string) (Tokens, error)
	DeleteByDevice(ctx context.Context, id int, ip, os, browser string) error
}
type Emails interface {
	SendUserVerificationEmail(VerificationEmailInput) error
	SendUserResetTokenEmail(ResetPasswordEmailInput) error
}
type Apartment interface {
	GetAllApartments(ctx context.Context, userId int) ([]*dto.GetApartmentResponse, error)
	GetApartmentById(ctx context.Context, userId int, id int) (*dto.GetApartmentResponse, error)
	CreateApartment(ctx context.Context, userId int, input dto.CreateApartmentInput) error
	DeleteApartment(ctx context.Context, userId int, id int) error
	UpdateApartment(ctx context.Context, userId int, id int, input *dto.UpdateApartmentInput) error
}
type Services struct {
	User
	Session
	Apartment
}

type Deps struct {
	Repos           *repository.Repository
	Hasher          hash.PasswordHasher
	StorageProvider storage.Provider
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	TokenManager    auth.TokenManager
	EmailSender     email.Sender
	EmailConfig     config.EmailConfig
	HTTPConfig      config.HTTPConfig
}
type VerificationEmailInput struct {
	Email            string
	Name             string
	VerificationCode string
}
type ResetPasswordEmailInput struct {
	Email      string
	Name       string
	ResetToken string
}

func NewServices(deps Deps) *Services {
	sessionService := NewSessionService(deps.Repos.Session, deps.AccessTokenTTL, deps.RefreshTokenTTL, deps.TokenManager)
	emailService := NewEmailService(deps.EmailSender, deps.EmailConfig, deps.HTTPConfig.BaseUrl)
	userService := NewUserService(
		deps.Repos.Users,
		deps.StorageProvider,
		deps.Hasher,
		deps.AccessTokenTTL,
		deps.RefreshTokenTTL,
		deps.TokenManager,
		sessionService,
		deps.EmailSender,
		emailService,
	)
	apartmentService := NewApartmentService(deps.Repos.Apartment)
	return &Services{
		User:      userService,
		Session:   sessionService,
		Apartment: apartmentService,
	}
}
