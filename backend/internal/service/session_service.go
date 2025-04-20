package service

import (
	"context"
	"strconv"
	"time"

	"github.com/scmbr/renting-app/internal/models"
	"github.com/scmbr/renting-app/internal/repository"
	"github.com/scmbr/renting-app/pkg/auth"
)

type SessionService struct {
	repo            repository.Session
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	tokenManager    auth.TokenManager
}

func NewSessionService(repo repository.Session, accessTTL, refreshTTL time.Duration, tokenManager auth.TokenManager) *SessionService {
	return &SessionService{
		repo:            repo,
		accessTokenTTL:  accessTTL,
		refreshTokenTTL: refreshTTL,
		tokenManager:    tokenManager,
	}
}

func (s *SessionService) CreateSession(ctx context.Context, userID int, ip string, os string, browser string) (Tokens, error) {
	var (
		res Tokens
		err error
	)
	res.AccessToken, err = s.tokenManager.NewJWT(strconv.Itoa(userID), s.accessTokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		return res, err
	}
	session := models.Session{
		UserID:       userID,
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(s.refreshTokenTTL),
		CreatedAt:    time.Now(),
		OS:           os,
		IP:           ip,
	}
	err = s.repo.CreateSession(ctx, session)
	return res, err
}
