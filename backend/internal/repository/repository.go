package repository

import (
	"context"
	"time"

	"github.com/scmbr/renting-app/internal/domain"
	"github.com/scmbr/renting-app/internal/dto"
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
	GetUser(email, password string) (domain.User, error)
	GetByCredentials(ctx context.Context, email, password string) (*dto.GetUser, error)
	Verify(ctx context.Context, code string) (dto.GetUser, error)
	GetByEmail(ctx context.Context, email string) (*dto.GetUser, error)
	UpdateVerificationCode(ctx context.Context, id int, verificationCode string) error
	SavePasswordResetToken(ctx context.Context, id int, resetToken string) error
	GetUserByResetToken(ctx context.Context, token string) (dto.GetUser, error)
	UpdatePasswordAndClearResetToken(ctx context.Context, userID int, newPassword string) error
	UpdateMe(input *dto.UpdateUser, userId int) (*dto.GetUser, error)
	UpdateRating(ctx context.Context, userID uint, rating float32) error
}
type Session interface {
	CreateSession(ctx context.Context, session domain.Session) error
	GetByRefreshToken(ctx context.Context, refreshToken string) (domain.Session, error)
	UpdateSession(ctx context.Context, session domain.Session) error
	GetByDevice(ctx context.Context, userID int, ip, os, browser string) (*domain.Session, error)
	UpdateTokens(ctx context.Context, sessionID int, refreshToken string, expiresAt time.Time) error
	DeleteByDevice(ctx context.Context, id int, ip, os, browser string) error
}
type Apartment interface {
	GetAllApartments(ctx context.Context, userId int) ([]*dto.GetApartmentResponse, error)
	GetApartmentById(ctx context.Context, userId int, id int) (*dto.GetApartmentResponse, error)
	CreateApartment(ctx context.Context, userId int, input dto.CreateApartmentInput) (uint, error)
	DeleteApartment(ctx context.Context, userId int, id int) error
	UpdateApartment(ctx context.Context, userId int, id int, input *dto.UpdateApartmentInput) error
	GetAllApartmentsAdmin(ctx context.Context) ([]*dto.GetApartmentResponse, error)
	GetApartmentByIdAdmin(ctx context.Context, id int) (*dto.GetApartmentResponse, error)
	UpdateApartmentAdmin(ctx context.Context, id int, input *dto.UpdateApartmentInput) error
	DeleteApartmentAdmin(ctx context.Context, id int) error
}
type Advert interface {
	GetAllUserAdverts(ctx context.Context, userId int) ([]*dto.GetAdvertResponse, error)
	GetUserAdvertById(ctx context.Context, userId int, id int) (*dto.GetAdvertResponse, error)
	CreateAdvert(ctx context.Context, userId int, input dto.CreateAdvertInput) error
	DeleteAdvert(ctx context.Context, userId int, id int) error
	UpdateAdvert(ctx context.Context, userId int, id int, input *dto.UpdateAdvertInput) error
	GetAllAdvertsAdmin(ctx context.Context) ([]*dto.GetAdvertResponse, error)
	GetAdvertByIdAdmin(ctx context.Context, id int) (*dto.GetAdvertResponse, error)
	UpdateAdvertAdmin(ctx context.Context, id int, input *dto.UpdateAdvertInput) error
	DeleteAdvertAdmin(ctx context.Context, id int) error
	GetAllAdverts(ctx context.Context, filter *dto.AdvertFilter) ([]domain.Advert, int64, error)
	GetAdvertById(ctx context.Context, id int) (*domain.Advert, error)
}
type ApartmentPhoto interface {
	GetAllPhotos(ctx context.Context, apartmentId int) ([]*dto.GetApartmentPhoto, error)
	GetPhotoById(ctx context.Context, apartmentId, photoId int) (*dto.GetApartmentPhoto, error)
	AddPhotoBatch(ctx context.Context, userId, apartmentId int, inputs []dto.CreatePhotoInput) error
	DeletePhoto(ctx context.Context, userId, apartmentId, photoId int) error
	SetCover(ctx context.Context, userId, apartmentId, photoId int) error
	HasCoverPhoto(apartmentId int) (bool, error)
	ReplaceAllPhotos(ctx context.Context, userId, apartmentId int, inputs []dto.CreatePhotoInput) error
}
type Favorites interface {
	GetAllFavorites(ctx context.Context, userId int) ([]dto.FavoriteResponseDTO, error)
	AddToFavorites(ctx context.Context, userId int, listingId int) error
	RemoveFromFavorites(ctx context.Context, userId int, listingId int) error
	IsFavorite(ctx context.Context, userId int, listingId int) (bool, error)
	GetUserFavorites(ctx context.Context, userID *int) (map[uint]bool, error)
}
type Notification interface {
	Save(notification dto.NotificationDTO) error
	GetByUserID(userID uint) ([]*dto.NotificationResponseDTO, error)
	MarkAsRead(notificationID uint) error
}
type Review interface {
	Create(ctx context.Context, authorID uint, input dto.CreateReviewInput) (*dto.GetReviewResponse, error)
	Update(ctx context.Context, reviewID uint, input dto.UpdateReviewInput) (*dto.GetReviewResponse, error)
	Delete(ctx context.Context, reviewID uint) error
	GetByAuthorID(ctx context.Context, authorID uint) ([]*dto.GetReviewResponse, error)
	GetByTargetID(ctx context.Context, targetID uint) ([]*dto.GetReviewResponse, error)
	ReviewExists(ctx context.Context, authorID uint, targetID uint) (bool, error)
	IsAuthor(ctx context.Context, userID uint, reviewID uint) (bool, error)
}
type Repository struct {
	Users
	Session
	Apartment
	Advert
	ApartmentPhoto
	Favorites
	Notification
	Review
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Users:          NewUsersRepo(db),
		Session:        NewSessionsRepo(db),
		Apartment:      NewApartmentRepo(db),
		Advert:         NewAdvertRepo(db),
		ApartmentPhoto: NewApartmentPhotoRepo(db),
		Favorites:      NewFavoritesRepo(db),
		Notification:   NewNotificationRepo(db),
		Review:         NewReviewRepo(db),
	}
}
