package service

import (
	"context"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/repository"
)

type FavoritesService struct {
	repo repository.Favorites
}

func NewFavoritesService(repo repository.Favorites) *FavoritesService {
	return &FavoritesService{
		repo: repo,
	}
}

func (s *FavoritesService) GetAllFavorites(ctx context.Context, userId int) ([]dto.FavoriteResponseDTO, error) {
	return s.repo.GetAllFavorites(ctx, userId)
}

func (s *FavoritesService) AddToFavorites(ctx context.Context, userId int, advertId int) error {
	return s.repo.AddToFavorites(ctx, userId, advertId)
}

func (s *FavoritesService) RemoveFromFavorites(ctx context.Context, userId int, advertId int) error {
	return s.repo.RemoveFromFavorites(ctx, userId, advertId)
}

func (s *FavoritesService) IsFavorite(ctx context.Context, userId int, advertId int) (bool, error) {
	return s.repo.IsFavorite(ctx, userId, advertId)
}
