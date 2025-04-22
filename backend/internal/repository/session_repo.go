package repository

import (
	"context"
	"time"

	"github.com/scmbr/renting-app/internal/models"
	"gorm.io/gorm"
)

type SessionsRepo struct {
	db *gorm.DB
}

func NewSessionsRepo(db *gorm.DB) *SessionsRepo {
	return &SessionsRepo{db: db}
}
func (r *SessionsRepo) CreateSession(ctx context.Context, session models.Session) error {
	return r.db.WithContext(ctx).Create(&session).Error
}
func (r *SessionsRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (models.Session, error) {
	var session models.Session
	result := r.db.Where("refresh_token = ?", refreshToken).First(&session)
	if result.Error != nil {
		return session, result.Error
	}

	return session, nil
}
func (r *SessionsRepo) UpdateSession(ctx context.Context, session models.Session) error {
	return r.db.WithContext(ctx).Save(&session).Error
}
func (r *SessionsRepo) UpdateTokens(ctx context.Context, sessionID int, refreshToken string, expiresAt time.Time) error {
	return r.db.WithContext(ctx).
		Model(&models.Session{}).
		Where("id = ?", sessionID).
		Updates(map[string]interface{}{
			"refresh_token": refreshToken,
			"expires_at":    expiresAt,
		}).Error
}
func (r *SessionsRepo) GetByDevice(ctx context.Context, userID int, ip, os, browser string) (*models.Session, error) {
	var session models.Session
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND ip = ? AND os = ? AND browser = ?", userID, ip, os, browser).
		First(&session).Error

	if err != nil {
		return nil, err
	}
	return &session, nil
}
