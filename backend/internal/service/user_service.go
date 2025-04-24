package service

import (
	"context"
	"crypto/rand"
	"math/big"
	"mime/multipart"
	"strings"
	"time"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/repository"
	"github.com/scmbr/renting-app/pkg/auth"
	"github.com/scmbr/renting-app/pkg/email"
	"github.com/scmbr/renting-app/pkg/hash"
	"github.com/scmbr/renting-app/pkg/storage"
	"github.com/sirupsen/logrus"
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
func generateVerificationCode() string {
	const codeLength = 6
	var code strings.Builder

	for i := 0; i < codeLength; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(10))
		code.WriteString(num.String())
	}

	return code.String()
}
func (s *UserService) SignUp(ctx context.Context, user dto.CreateUser) error {
	passwordHash, err := s.hasher.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = passwordHash
	verificationCode := generateVerificationCode()
	if err := s.repo.CreateUser(user, verificationCode); err != nil {
		return err
	}
	return s.emailService.SendUserVerificationEmail(VerificationEmailInput{
		Email:            user.Email,
		Name:             user.Name,
		VerificationCode: verificationCode,
	})
}
func (s *UserService) SignIn(ctx context.Context, email string, password string, ip string, os string, browser string) (Tokens, error) {
	passwordHash, err := s.hasher.Hash(password)
	if err != nil {
		return Tokens{}, err
	}

	user, err := s.repo.GetByCredentials(ctx, email, passwordHash)
	if err != nil {
		return Tokens{}, err
	}

	return s.sessionService.CreateSession(ctx, user.Id, ip, os, browser)
}
func (s *UserService) VerifyEmail(ctx context.Context, code string) error {
	user, err := s.repo.Verify(ctx, code)
	if err != nil {

		return err
	}

	logrus.Info(user)
	return nil
}
