package service

import (
	"context"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/repository"
)

type AdvertService struct {
	advertRepo         repository.Advert
	favoritesRepo      repository.Favorites
	apartmentPhotoRepo repository.ApartmentPhoto
}

func NewAdvertService(advertRepo repository.Advert, favoritesRepo repository.Favorites, apartmentRepo repository.ApartmentPhoto) *AdvertService {
	return &AdvertService{
		advertRepo:         advertRepo,
		favoritesRepo:      favoritesRepo,
		apartmentPhotoRepo: apartmentRepo,
	}
}

func (s *AdvertService) GetAllAdverts(ctx context.Context, userID *int, filter *dto.AdvertFilter) ([]*dto.GetAdvertResponse, int64, error) {
	adverts, total, err := s.advertRepo.GetAllAdverts(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	var favoriteMap map[uint]bool
	if userID != nil {
		favoriteMap, err = s.favoritesRepo.GetUserFavorites(ctx, userID)
		if err != nil {
			return nil, 0, err
		}
	}
	result := make([]*dto.GetAdvertResponse, len(adverts))
	for i, adv := range adverts {
		resp := dto.FromAdvert(adv)
		if userID != nil {
			resp.IsFavorite = favoriteMap[adv.ID]
		} else {
			resp.IsFavorite = false
		}
		photosPtrs, err := s.apartmentPhotoRepo.GetAllPhotos(ctx, int(adv.ApartmentID))
		if err != nil {
			return nil, 0, err
		}
		photos := make([]dto.GetApartmentPhoto, len(photosPtrs))
		for j, p := range photosPtrs {
			photos[j] = *p
		}
		resp.Apartment.ApartmentPhotos = photos
		result[i] = resp
	}

	return result, total, nil
}

func (s *AdvertService) GetAdvertById(ctx context.Context, id int) (*dto.GetAdvertResponse, error) {
	return s.advertRepo.GetAdvertById(ctx, id)
}

func (s *AdvertService) GetAllUserAdverts(ctx context.Context, userID int) ([]*dto.GetAdvertResponse, error) {
	return s.advertRepo.GetAllUserAdverts(ctx, userID)
}

func (s *AdvertService) GetUserAdvertById(ctx context.Context, userID int, id int) (*dto.GetAdvertResponse, error) {
	return s.advertRepo.GetUserAdvertById(ctx, userID, id)
}

func (s *AdvertService) CreateAdvert(ctx context.Context, userID int, input dto.CreateAdvertInput) error {
	return s.advertRepo.CreateAdvert(ctx, userID, input)
}

func (s *AdvertService) DeleteAdvert(ctx context.Context, userID int, id int) error {
	return s.advertRepo.DeleteAdvert(ctx, userID, id)
}

func (s *AdvertService) UpdateAdvert(ctx context.Context, userID int, id int, input *dto.UpdateAdvertInput) error {
	return s.advertRepo.UpdateAdvert(ctx, userID, id, input)
}

func (s *AdvertService) GetAllAdvertsAdmin(ctx context.Context) ([]*dto.GetAdvertResponse, error) {
	return s.advertRepo.GetAllAdvertsAdmin(ctx)
}

func (s *AdvertService) GetAdvertByIdAdmin(ctx context.Context, id int) (*dto.GetAdvertResponse, error) {
	return s.advertRepo.GetAdvertByIdAdmin(ctx, id)
}

func (s *AdvertService) UpdateAdvertAdmin(ctx context.Context, id int, input *dto.UpdateAdvertInput) error {
	return s.advertRepo.UpdateAdvertAdmin(ctx, id, input)
}

func (s *AdvertService) DeleteAdvertAdmin(ctx context.Context, id int) error {
	return s.advertRepo.DeleteAdvertAdmin(ctx, id)
}
