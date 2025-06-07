package service

import (
	"context"
	"fmt"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/repository"
	"github.com/sirupsen/logrus"
)

type FavoritesService struct {
	FavoritesRepo       repository.Favorites
	UserRepo            repository.Users
	AdvertRepo          repository.Advert
	notificationService *NotificationService
}

func NewFavoritesService(
	favoritesRepo repository.Favorites,
	userRepo repository.Users,
	advertRepo repository.Advert,
	notificationService *NotificationService,
) *FavoritesService {
	return &FavoritesService{
		FavoritesRepo:       favoritesRepo,
		UserRepo:            userRepo,
		AdvertRepo:          advertRepo,
		notificationService: notificationService,
	}
}

func (s *FavoritesService) GetAllFavorites(ctx context.Context, userId int) ([]dto.FavoriteResponseDTO, error) {
	return s.FavoritesRepo.GetAllFavorites(ctx, userId)
}

func (s *FavoritesService) AddToFavorites(ctx context.Context, userId int, advertId int) error {
	err := s.FavoritesRepo.AddToFavorites(ctx, userId, advertId)
	if err != nil {
		return err
	}
	user, err := s.UserRepo.GetUserById(userId)
	if err != nil {
		return err
	}
	advert, err := s.AdvertRepo.GetAdvertById(ctx, advertId)
	if err != nil {
		return err
	}
	ownerID := advert.UserID
	notification := dto.NotificationDTO{
		UserID:   ownerID,
		Type:     "Избранное",
		Title:    "Ваше объявление кому-то понравилось",
		AdvertId: uint(advertId),
		Content:  fmt.Sprintf("%s добавил/а ваше объявление в избранное", user.Name),
	}
	err = s.notificationService.CreateAndSend(notification)
	if err != nil {
		logrus.Warnf("не удалось отправить уведомление: %v", err)
	}
	return nil
}

func (s *FavoritesService) RemoveFromFavorites(ctx context.Context, userId int, advertId int) error {
	return s.FavoritesRepo.RemoveFromFavorites(ctx, userId, advertId)
}

func (s *FavoritesService) IsFavorite(ctx context.Context, userId int, advertId int) (bool, error) {
	return s.FavoritesRepo.IsFavorite(ctx, userId, advertId)
}
