package service

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/repository"
	"github.com/scmbr/renting-app/pkg/auth"
	"github.com/scmbr/renting-app/pkg/email"
	"github.com/scmbr/renting-app/pkg/hash"
	"github.com/scmbr/renting-app/pkg/storage"
)

type UserService struct {
	repo            repository.Users
	storage         storage.Provider
	hasher          hash.PasswordHasher
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	tokenManager    auth.TokenManager
	sessionService  Session
	emailService    Emails
	smtp            email.Sender
}

func NewUserService(repo repository.Users, storage storage.Provider, hasher hash.PasswordHasher, accessTTL, refreshTTL time.Duration, tokenManager auth.TokenManager, sessionService Session, smtp email.Sender, emailService Emails) *UserService {
	return &UserService{
		repo:            repo,
		storage:         storage,
		hasher:          hasher,
		tokenManager:    tokenManager,
		accessTokenTTL:  accessTTL,
		refreshTokenTTL: refreshTTL,
		sessionService:  sessionService,
		smtp:            smtp,
		emailService:    emailService,
	}
}

func (s *UserService) GetAllUsers() ([]dto.GetUser, error) {

	return s.repo.GetAllUsers()
}
func (s *UserService) GetUserById(id int) (*dto.GetUser, error) {

	return s.repo.GetUserById(id)
}
func (s *UserService) DeleteUserById(id int) (*dto.GetUser, error) {

	return s.repo.DeleteUserById(id)
}
func (s *UserService) UpdateUserById(input *dto.GetUser) (*dto.GetUser, error) {
	return s.repo.UpdateUserById(input)
}
func (s *UserService) UploadAvatarToS3(ctx context.Context, fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()
	return s.storage.Upload(ctx, storage.UploadInput{
		File:        file,
		Name:        fileHeader.Filename,
		Size:        fileHeader.Size,
		ContentType: "image/png",
	})

}
func (s *UserService) UpdateAvatar(userId int, avatarURL string) error {

	return s.repo.UpdateAvatar(userId, avatarURL)
}
