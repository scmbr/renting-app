package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/sirupsen/logrus"
)

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

	return s.sessionService.CreateSession(ctx, user.Role, user.Id, ip, os, browser)
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

func generateVerificationCode() string {
	const codeLength = 6
	var code strings.Builder

	for i := 0; i < codeLength; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(10))
		code.WriteString(num.String())
	}

	return code.String()
}

func generateResetToken() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatalf("failed to generate reset token: %v", err)
	}
	return hex.EncodeToString(bytes)
}
