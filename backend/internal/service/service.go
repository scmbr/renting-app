package service

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/scmbr/renting-app/internal/config"
	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/infrastructure/redis/cache"
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
	GetAllUsers() ([]dto.GetUserResponse, error)
	GetUserById(id int) (*dto.GetProfileResponse, error)
	DeleteUserById(id int) (*dto.GetUserResponse, error)
	UpdateUserById(input *dto.GetUserResponse) (*dto.GetUserResponse, error)
	// UploadAvatarToS3(ctx context.Context, fileHeader *multipart.FileHeader) (string, error)
	UpdateAvatar(userId int, avatarURL string) error
	SignIn(ctx context.Context, email string, password string, ip string, os string, browser string) (Tokens, error)
	SignUp(ctx context.Context, user dto.CreateUser) error
	VerifyEmail(ctx context.Context, code string) (*dto.GetUserResponse, error)
	ResendVerificationCode(ctx context.Context, email string) error
	ForgotPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, resetToken string, newPassword string) error
	LogOut(ctx context.Context, id int, ip, os, browser string) error
	GenerateTokens(ctx context.Context, email string, ip string, os string, browser string) (Tokens, error)
	UpdateMe(userID int, input dto.UpdateUser) error
}
type Session interface {
	CreateSession(ctx context.Context, role string, userID int, ip string, os string, browser string) (Tokens, error)
	RefreshSession(ctx context.Context, role string, refreshToken, ip, os, browser string) (Tokens, error)
	DeleteByDevice(ctx context.Context, id int, ip, os, browser string) error
}
type Emails interface {
	SendUserVerificationEmail(VerificationEmailInput) error
	SendUserResetTokenEmail(ResetPasswordEmailInput) error
}
type Apartment interface {
	GetAllApartments(ctx context.Context, userId int) ([]*dto.GetApartmentResponse, error)
	GetApartmentById(ctx context.Context, userId int, id int) (*dto.GetApartmentResponse, error)
	CreateApartment(ctx context.Context, userId int, input dto.CreateApartmentInput) (*dto.GetApartmentResponse, error)
	DeleteApartment(ctx context.Context, userId int, id int) error
	UpdateApartment(ctx context.Context, userId int, id int, input *dto.UpdateApartmentInput) error
	GetAllApartmentsAdmin(ctx context.Context) ([]*dto.GetApartmentResponse, error)
	GetApartmentByIdAdmin(ctx context.Context, id int) (*dto.GetApartmentResponse, error)
	UpdateApartmentAdmin(ctx context.Context, id int, input *dto.UpdateApartmentInput) error
	DeleteApartmentAdmin(ctx context.Context, id int) error
}
type Advert interface {
	GetAllUserAdverts(ctx context.Context, userId int, filter *dto.AdvertFilter) ([]*dto.GetAdvertResponse, int64, error)
	GetUserAdvertById(ctx context.Context, userId int, id int) (*dto.GetAdvertResponse, error)
	CreateAdvert(ctx context.Context, userId int, input dto.CreateAdvertInput) (*dto.GetAdvertResponse, error)
	DeleteAdvert(ctx context.Context, userId int, id int) error
	UpdateAdvert(ctx context.Context, userId int, id int, input *dto.UpdateAdvertInput) error
	GetAllAdvertsAdmin(ctx context.Context, filter *dto.AdvertFilter) ([]*dto.GetAdvertResponse, int64, error)
	GetAdvertByIdAdmin(ctx context.Context, id int) (*dto.GetAdvertResponse, error)
	UpdateAdvertAdmin(ctx context.Context, id int, input *dto.UpdateAdvertInput) error
	DeleteAdvertAdmin(ctx context.Context, id int) error
	GetAllAdverts(ctx context.Context, userId *int, filter *dto.AdvertFilter) ([]*dto.GetAdvertResponse, int64, error)
	GetAdvertById(ctx context.Context, id int) (*dto.GetAdvertResponse, error)
}

type ApartmentPhoto interface {
	GetAllPhotos(ctx context.Context, apartmentId int) ([]dto.GetApartmentPhotoResponse, error)
	GetPhotoById(ctx context.Context, apartmentId, photoId int) (*dto.GetApartmentPhotoResponse, error)
	AddPhotos(ctx context.Context, userId, apartmentId int, files []*multipart.FileHeader) ([]*dto.GetApartmentPhotoResponse, error)
	DeletePhoto(ctx context.Context, userId, apartmentId, photoId int) error
	SetCover(ctx context.Context, userId, apartmentId, photoId int) error
	UploadPhotoToS3(ctx context.Context, fileHeader *multipart.FileHeader) (string, error)
	HasCoverPhoto(ctx context.Context, apartmentId int) (bool, error)
	ReplaceAllPhotos(ctx context.Context, userId, apartmentId int, inputs []dto.CreatePhotoInput) error
}
type Favorites interface {
	GetAllFavorites(ctx context.Context, userId int) ([]dto.FavoriteResponseDTO, error)
	AddToFavorites(ctx context.Context, userId int, advertId int) error
	RemoveFromFavorites(ctx context.Context, userId int, advertId int) error
	IsFavorite(ctx context.Context, userId int, advertId int) (bool, error)
}
type Notification interface {
	CreateAndSend(notification dto.NotificationDTO) error
	GetUserNotifications(userID uint) ([]*dto.NotificationResponseDTO, error)
	MarkAsRead(notificationID uint) error
}
type Review interface {
	Create(ctx context.Context, authorID uint, input dto.CreateReviewInput) (*dto.GetReviewResponse, error)
	Update(ctx context.Context, userID uint, reviewID uint, input dto.UpdateReviewInput) (*dto.GetReviewResponse, error)
	Delete(ctx context.Context, userID uint, reviewID uint) error
	GetByAuthorID(ctx context.Context, authorID uint) ([]*dto.GetReviewResponse, error)
	GetByTargetID(ctx context.Context, targetID uint) ([]*dto.GetReviewResponse, error)
}
type Services struct {
	User
	Session
	Apartment
	Advert
	ApartmentPhoto
	Favorites
	Notification
	Review
}

type Deps struct {
	Repos              *repository.Repository
	Hasher             hash.PasswordHasher
	StorageProvider    storage.Provider
	AccessTokenTTL     time.Duration
	RefreshTokenTTL    time.Duration
	TokenManager       auth.TokenManager
	EmailSender        email.Sender
	EmailConfig        config.EmailConfig
	HTTPConfig         config.HTTPConfig
	NotificationSender NotificationSender
	Cache              cache.Cache
	CacheTTL           time.Duration
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
	apartmentService := NewApartmentService(deps.Repos.Apartment,deps.Repos.ApartmentPhoto)
	apartmentPhotoService := NewApartmentPhotoService(deps.Repos.ApartmentPhoto, deps.StorageProvider)
	advertService := NewAdvertService(deps.Repos.Advert, deps.Repos.Favorites, deps.Repos.ApartmentPhoto, deps.Cache, deps.CacheTTL)

	notificationService := NewNotificationService(deps.Repos.Notification, deps.NotificationSender)
	favoritesService := NewFavoritesService(deps.Repos.Favorites, deps.Repos.Users, deps.Repos.Advert, notificationService)
	reviewService := NewReviewService(deps.Repos.Review, deps.Repos.Users)
	return &Services{
		User:           userService,
		Session:        sessionService,
		Apartment:      apartmentService,
		Advert:         advertService,
		ApartmentPhoto: apartmentPhotoService,
		Favorites:      favoritesService,
		Notification:   notificationService,
		Review:         reviewService,
	}
}
