package service

import (
	"context"
	"fmt"
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

func (s *SessionService) CreateSession(ctx context.Context, role string, userID int, ip string, os string, browser string) (Tokens, error) {
	var (
		res Tokens
		err error
	)
	res.AccessToken, err = s.tokenManager.NewJWT(role, strconv.Itoa(userID), s.accessTokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = s.tokenManager.NewJWT(role, strconv.Itoa(userID), s.refreshTokenTTL)
	if err != nil {
		return res, err
	}
	existingSession, err := s.repo.GetByDevice(ctx, userID, ip, os, browser)
	if err == nil && existingSession != nil {
		err = s.repo.UpdateTokens(ctx, existingSession.ID, res.RefreshToken, time.Now().Add(s.refreshTokenTTL))
		return res, err
	}
	session := models.Session{
		UserID:       userID,
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(s.refreshTokenTTL),
		CreatedAt:    time.Now(),
		OS:           os,
		IP:           ip,
		Browser:      browser,
	}
	err = s.repo.CreateSession(ctx, session)
	return res, err
}
func (s *SessionService) RefreshSession(ctx context.Context, role string, refreshToken, ip, os, browser string) (Tokens, error) {
	session, err := s.repo.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		return Tokens{}, err
	}
	// Генерация новых токенов
	var res Tokens
	res.AccessToken, err = s.tokenManager.NewJWT(role, strconv.Itoa(session.UserID), s.accessTokenTTL)
	if err != nil {
		return res, err
	}
	res.RefreshToken, err = s.tokenManager.NewJWT(role, strconv.Itoa(session.UserID), s.refreshTokenTTL)
	if err != nil {
		return res, err
	}

	// Обновление сессии
	session.RefreshToken = res.RefreshToken
	session.ExpiresAt = time.Now().Add(s.refreshTokenTTL)
	session.IP = ip
	session.OS = os
	session.Browser = browser

	err = s.repo.UpdateSession(ctx, session)
	if err != nil {
		return Tokens{}, fmt.Errorf("failed to update session: %w", err)
	}

	return res, nil
}
func (s *SessionService) DeleteByDevice(ctx context.Context, id int, ip, os, browser string) error {
	err := s.repo.DeleteByDevice(ctx, id, ip, os, browser)
	if err != nil {
		return err
	}
	return nil
}
