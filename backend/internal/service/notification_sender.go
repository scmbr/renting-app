package service

import "github.com/scmbr/renting-app/internal/dto"

type NotificationSender interface {
    SendNotification(userID uint, notification dto.NotificationResponseDTO) error
}
