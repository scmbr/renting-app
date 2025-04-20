package repository

import (
	"context"

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
