package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/repository"
	"github.com/scmbr/renting-app/pkg/hash"
)

const (
	salt       = "sadfasdw234;fgd9"
	signingKey = "asdsadasdasd"
	tokenTTL   = 14 * time.Hour * 24
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
type AuthService struct {
	repo   repository.Authorization
	hasher hash.PasswordHasher
}

func NewAuthService(repo repository.Authorization, hasher hash.PasswordHasher) *AuthService {
	return &AuthService{
		repo:   repo,
		hasher: hasher,
	}
}

func (s *AuthService) CreateUser(user dto.CreateUser) (int, error) {
	passwordHash, err := s.hasher.Hash(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = passwordHash
	return s.repo.CreateUser(user)
}
func (s *AuthService) GenerateToken(email string, password string) (string, error) {
	passwordHash, err := s.hasher.Hash(password)
	if err != nil {
		return "", err
	}
	user, err := s.repo.GetUser(email, passwordHash)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		int(user.ID),
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}
