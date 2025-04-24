package service

import (
	"fmt"

	"github.com/scmbr/renting-app/internal/config"
	"github.com/scmbr/renting-app/pkg/email"
)

type EmailService struct {
	sender email.Sender
	config config.EmailConfig
}
type verificationEmailInput struct {
	Code string
}

func NewEmailService(sender email.Sender, config config.EmailConfig) *EmailService {
	return &EmailService{
		sender: sender,
		config: config,
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
