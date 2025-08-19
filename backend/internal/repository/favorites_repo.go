package repository

import (
	"context"
	"errors"

	"github.com/scmbr/renting-app/internal/domain"
	"github.com/scmbr/renting-app/internal/dto"
	"gorm.io/gorm"
)

type FavoritesRepo struct {
	db *gorm.DB
}

func NewFavoritesRepo(db *gorm.DB) *FavoritesRepo {
	return &FavoritesRepo{db: db}
}

func (r *FavoritesRepo) GetAllFavorites(ctx context.Context, userId int) ([]dto.FavoriteResponseDTO, error) {
	var favorites []domain.Favorites

	err := r.db.WithContext(ctx).
		Where("user_id = ?", userId).
		Find(&favorites).Error
	if err != nil {
		return nil, err
	}

	var result []dto.FavoriteResponseDTO
	for _, fav := range favorites {
		result = append(result, dto.FavoriteResponseDTO{
			ID:        fav.ID,
			AdvertID:  fav.AdvertID,
			CreatedAt: fav.CreatedAt,
		})
	}

	return result, nil
}
func (r *FavoritesRepo) AddToFavorites(ctx context.Context, userId int, advertId int) error {

	var count int64
	err := r.db.WithContext(ctx).
		Model(&domain.Favorites{}).
		Where("user_id = ? AND advert_id = ?", userId, advertId).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {

		return errors.New("already in favorites")
	}

	fav := domain.Favorites{
		UserID:   uint(userId),
		AdvertID: uint(advertId),
	}

	if err := r.db.WithContext(ctx).Create(&fav).Error; err != nil {
		return err
	}
	return nil
}

func (r *FavoritesRepo) RemoveFromFavorites(ctx context.Context, userId int, advertId int) error {
	result := r.db.WithContext(ctx).
		Where("user_id = ? AND advert_id = ?", userId, advertId).
		Delete(&domain.Favorites{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("favorite not found")
	}
	return nil
}

func (r *FavoritesRepo) IsFavorite(ctx context.Context, userId int, advertId int) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&domain.Favorites{}).
		Where("user_id = ? AND advert_id = ?", userId, advertId).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func (r *FavoritesRepo) GetUserFavorites(ctx context.Context, userID *int) (map[uint]bool, error) {
	var favorites []domain.Favorites
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&favorites).Error
	if err != nil {
		return nil, err
	}

	likedMap := make(map[uint]bool, len(favorites))
	for _, f := range favorites {
		likedMap[f.AdvertID] = true
	}
	return likedMap, nil
}
