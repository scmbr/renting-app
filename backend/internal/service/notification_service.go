package service

import (
	"time"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/repository"
)

type NotificationService struct {
	repo   repository.Notification
	sender NotificationSender
}

func NewNotificationService(repo repository.Notification, sender NotificationSender) *NotificationService {
	return &NotificationService{
		repo:   repo,
		sender: sender,
	}
}

func (s *NotificationService) CreateAndSend(notification dto.NotificationDTO) error {
	err := s.repo.Save(notification)
	if err != nil {
		return err
	}

	response := dto.NotificationResponseDTO{
		Type:      notification.Type,
		Title:     notification.Title,
		Content:   notification.Content,
		AdvertId:  notification.AdvertId,
		IsRead:    false,
		CreatedAt: time.Now(),
	}

	_ = s.sender.SendNotification(notification.UserID, response)

	return nil
}

func (s *NotificationService) GetUserNotifications(userID uint) ([]*dto.NotificationResponseDTO, error) {
	return s.repo.GetByUserID(userID)
}

func (s *NotificationService) MarkAsRead(notificationID uint) error {
	return s.repo.MarkAsRead(notificationID)
}
