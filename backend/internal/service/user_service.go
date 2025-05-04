package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
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
	if err := s.repo.CreateUser(ctx, user, verificationCode); err != nil {
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
func (s *UserService) ResendVerificationCode(ctx context.Context, email string) error {
	verificationCode := generateVerificationCode()
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return err
	}
	if user.Verified {
		return errors.New("user already verified")
	}
	err = s.repo.UpdateVerificationCode(ctx, user.Id, verificationCode)
	if err != nil {
		return err
	}
	return s.emailService.SendUserVerificationEmail(VerificationEmailInput{
		Email:            user.Email,
		Name:             user.Name,
		VerificationCode: verificationCode,
	})
}
func (s *UserService) ForgotPassword(ctx context.Context, email string) error {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	// Генерируем токен для сброса пароля
	resetToken := generateResetToken()

	err = s.repo.SavePasswordResetToken(ctx, user.Id, resetToken)
	if err != nil {
		return fmt.Errorf("failed to save reset token: %w", err)
	}

	return s.emailService.SendUserResetTokenEmail(ResetPasswordEmailInput{
		Email:      user.Email,
		Name:       user.Name,
		ResetToken: resetToken,
	})
}
func generateResetToken() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatalf("failed to generate reset token: %v", err)
	}
	return hex.EncodeToString(bytes)
}
func (s *UserService) ResetPassword(ctx context.Context, resetToken string, newPassword string) error {
	user, err := s.repo.GetUserByResetToken(ctx, resetToken)
	if err != nil {
		return fmt.Errorf("invalid or expired reset token: %w", err)
	}

	hashedPassword, err := s.hasher.Hash(newPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	err = s.repo.UpdatePasswordAndClearResetToken(ctx, user.Id, hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}
func (s *UserService) LogOut(ctx context.Context, id int, ip, os, browser string) error {
	err := s.sessionService.DeleteByDevice(ctx, id, ip, os, browser)
	if err != nil {
		return fmt.Errorf("failed to log out: %w", err)
	}

	return nil
}
