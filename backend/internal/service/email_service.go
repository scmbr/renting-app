package service

import (
	"fmt"

	"github.com/scmbr/renting-app/internal/config"
	"github.com/scmbr/renting-app/pkg/email"
)

type EmailService struct {
	sender  email.Sender
	config  config.EmailConfig
	baseURL string
}
type verificationEmailInput struct {
	Code string
}
type resetPasswrodInput struct {
	Link string
}

func NewEmailService(sender email.Sender, config config.EmailConfig, baseURL string) *EmailService {
	return &EmailService{
		sender:  sender,
		config:  config,
		baseURL: baseURL,
	}
}
func (s *EmailService) SendUserVerificationEmail(input VerificationEmailInput) error {
	subject := fmt.Sprintf("Спасибо за регистрацию, %s!", input.Name)

	sendInput := email.SendEmailInput{Subject: subject, To: input.Email}
	templateInput := verificationEmailInput{Code: input.VerificationCode}
	if err := sendInput.GenerateBodyFromHTML(s.config.Templates.Verification, templateInput); err != nil {
		return err
	}

	return s.sender.Send(sendInput)
}
func (s *EmailService) SendUserResetTokenEmail(input ResetPasswordEmailInput) error {
	subject := "Сброс пароля"

	sendInput := email.SendEmailInput{Subject: subject, To: input.Email}
	resetPasswordLink := fmt.Sprintf("%s/reset-password?token=%s", s.baseURL, input.ResetToken)
	templateInput := resetPasswrodInput{Link: resetPasswordLink}
	if err := sendInput.GenerateBodyFromHTML(s.config.Templates.Reset, templateInput); err != nil {
		return err
	}

	return s.sender.Send(sendInput)
}
