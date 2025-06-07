package repository

import (
	"time"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/models"
	"gorm.io/gorm"
)

type NotificationRepo struct {
	db *gorm.DB
}

func NewNotificationRepo(db *gorm.DB) *NotificationRepo {
	return &NotificationRepo{db: db}
}

func (r *NotificationRepo) Save(notification dto.NotificationDTO) error {
	notif := models.Notification{
		UserID:    notification.UserID,
		Type:      notification.Type,
		Title:     notification.Title,
		Content:   notification.Content,
		AdvertID:  notification.AdvertId,
		IsSent:    false,
		IsRead:    false,
		CreatedAt: time.Now(),
	}
	return r.db.Create(&notif).Error
}
func (r *NotificationRepo) GetByUserID(userID uint) ([]*dto.NotificationResponseDTO, error) {
	var notifications []models.Notification

	err := r.db.
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&notifications).Error

	if err != nil {
		return nil, err
	}

	result := make([]*dto.NotificationResponseDTO, len(notifications))
	for i, n := range notifications {
		result[i] = &dto.NotificationResponseDTO{
			ID:        n.ID,
			Type:      n.Type,
			Title:     n.Title,
			AdvertId:  n.AdvertID,
			Content:   n.Content,
			IsRead:    n.IsRead,
			CreatedAt: n.CreatedAt,
		}
	}

	return result, nil
}
func (r *NotificationRepo) MarkAsRead(notificationID uint) error {
	return r.db.Model(&models.Notification{}).
		Where("id = ?", notificationID).
		Update("is_read", true).Error
}
