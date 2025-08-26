package service

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/infrastructure/redis/cache"
	"github.com/scmbr/renting-app/internal/repository"
)

type AdvertService struct {
	advertRepo         repository.Advert
	favoritesRepo      repository.Favorites
	apartmentPhotoRepo repository.ApartmentPhoto
	cache              cache.Cache
	cacheTTL           time.Duration
}

func NewAdvertService(advertRepo repository.Advert, favoritesRepo repository.Favorites, apartmentRepo repository.ApartmentPhoto, cache cache.Cache, cacheTTL time.Duration) *AdvertService {
	return &AdvertService{
		advertRepo:         advertRepo,
		favoritesRepo:      favoritesRepo,
		apartmentPhotoRepo: apartmentRepo,
		cache:              cache,
		cacheTTL:           cacheTTL,
	}
}

func (s *AdvertService) GetAllAdverts(ctx context.Context, userID *int, filter *dto.AdvertFilter) ([]*dto.GetAdvertResponse, int64, error) {
	filterJSON, err := json.Marshal(filter)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to marshal filter: %w", err)
	}
	hash := sha256.Sum256(filterJSON)
	cacheKey := fmt.Sprintf("adverts:%x", hash)
	if cached, err := s.cache.Get(ctx, cacheKey); err == nil && cached != "" {
		var resp struct {
			Adverts []*dto.GetAdvertResponse
			Total   int64
		}
		if jsonErr := json.Unmarshal([]byte(cached), &resp); jsonErr == nil {
			if userID != nil {
				favoriteMap, err := s.favoritesRepo.GetUserFavorites(ctx, userID)
				if err != nil {
					return nil, 0, err
				}
				for _, adv := range resp.Adverts {
					_, adv.IsFavorite = favoriteMap[adv.ID]
				}
			}
			return resp.Adverts, resp.Total, nil
		} else {
			log.Printf("failed to unmarshal cached data for key %s: %v", cacheKey, jsonErr)
		}
	}

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
	apartmentIDs := make([]uint, len(adverts))
	for i, adv := range adverts {
		apartmentIDs[i] = adv.ApartmentID
	}
	photosMap, err := s.apartmentPhotoRepo.GetAllPhotosForApartments(ctx, apartmentIDs)
	if err != nil {
		return nil, 0, err
	}
	for i, adv := range adverts {
		resp := dto.FromAdvert(adv)
		if userID != nil {
			resp.IsFavorite = favoriteMap[adv.ID]
		} else {
			resp.IsFavorite = false
		}

		resp.Apartment.ApartmentPhotos = photosMap[adv.ApartmentID]
		if resp.Apartment.ApartmentPhotos == nil {
			resp.Apartment.ApartmentPhotos = []dto.GetApartmentPhotoResponse{}
		}
		result[i] = resp
	}

	payload, err := json.Marshal(struct {
		Adverts []*dto.GetAdvertResponse
		Total   int64
	}{result, total})
	if err != nil {
		log.Printf("failed to marshal payload: %v", err)
	} else if err := s.cache.Set(ctx, cacheKey, string(payload), s.cacheTTL); err != nil {
		log.Printf("failed to set cache: %v", err)
	}

	return result, total, nil
}

func (s *AdvertService) GetAdvertById(ctx context.Context, id int) (*dto.GetAdvertResponse, error) {
	advert, err := s.advertRepo.GetAdvertById(ctx, id)
	if err != nil {
		return nil, err
	}
	resp := dto.FromAdvert(advert)
	photos, err := s.apartmentPhotoRepo.GetAllPhotos(ctx, int(advert.ApartmentID))
	if err != nil {
		return nil, err
	}
	resp.Apartment.ApartmentPhotos = photos
	return resp, nil
}

func (s *AdvertService) GetAllUserAdverts(ctx context.Context, userId int, filter *dto.AdvertFilter) ([]*dto.GetAdvertResponse, int64, error) {
	filter.UserID = uint(userId)

	adverts, total, err := s.advertRepo.GetAllAdverts(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	var favoriteMap map[uint]bool

	favoriteMap, err = s.favoritesRepo.GetUserFavorites(ctx, &userId)
	if err != nil {
		return nil, 0, err
	}

	result := make([]*dto.GetAdvertResponse, len(adverts))
	apartmentIDs := make([]uint, len(adverts))
	for i, adv := range adverts {
		apartmentIDs[i] = adv.ApartmentID
	}

	photosMap, err := s.apartmentPhotoRepo.GetAllPhotosForApartments(ctx, apartmentIDs)
	if err != nil {
		return nil, 0, err
	}
	for i, adv := range adverts {
		resp := dto.FromAdvert(adv)

		resp.IsFavorite = favoriteMap[adv.ID]

		resp.Apartment.ApartmentPhotos = photosMap[adv.ApartmentID]
		result[i] = resp
	}
	return result, total, nil
}

func (s *AdvertService) GetUserAdvertById(ctx context.Context, userID int, id int) (*dto.GetAdvertResponse, error) {
	advert, err := s.advertRepo.GetAdvertById(ctx, id)
	if err != nil {
		return nil, err
	}
	resp := dto.FromAdvert(advert)
	photos, err := s.apartmentPhotoRepo.GetAllPhotos(ctx, int(advert.ApartmentID))
	if err != nil {
		return nil, err
	}
	resp.Apartment.ApartmentPhotos = photos
	return resp, nil
}

func (s *AdvertService) CreateAdvert(ctx context.Context, userID int, input dto.CreateAdvertInput) (*dto.GetAdvertResponse, error) {
	advert, err := s.advertRepo.CreateAdvert(ctx, userID, input)
	if err != nil {
		return nil, err
	}
	resp := dto.FromAdvert(advert)
	return resp, nil
}

func (s *AdvertService) DeleteAdvert(ctx context.Context, userID int, id int) error {
	return s.advertRepo.DeleteAdvert(ctx, id)
}

func (s *AdvertService) UpdateAdvert(ctx context.Context, userID int, id int, input *dto.UpdateAdvertInput) error {
	return s.advertRepo.UpdateAdvert(ctx, id, input)
}

func (s *AdvertService) GetAllAdvertsAdmin(ctx context.Context, filter *dto.AdvertFilter) ([]*dto.GetAdvertResponse, int64, error) {
	adverts, total, err := s.advertRepo.GetAllAdverts(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	result := make([]*dto.GetAdvertResponse, len(adverts))
	for i, adv := range adverts {
		resp := dto.FromAdvert(adv)

		photos, err := s.apartmentPhotoRepo.GetAllPhotos(ctx, int(adv.ApartmentID))
		if err != nil {
			return nil, 0, err
		}

		resp.Apartment.ApartmentPhotos = photos
		result[i] = resp
	}
	return result, total, err
}
func (s *AdvertService) GetAdvertByIdAdmin(ctx context.Context, id int) (*dto.GetAdvertResponse, error) {
	advert, err := s.advertRepo.GetAdvertById(ctx, id)
	if err != nil {
		return nil, err
	}
	resp := dto.FromAdvert(advert)
	photos, err := s.apartmentPhotoRepo.GetAllPhotos(ctx, int(advert.ApartmentID))
	if err != nil {
		return nil, err
	}
	resp.Apartment.ApartmentPhotos = photos
	return resp, nil
}

func (s *AdvertService) UpdateAdvertAdmin(ctx context.Context, id int, input *dto.UpdateAdvertInput) error {
	return s.advertRepo.UpdateAdvert(ctx, id, input)
}

func (s *AdvertService) DeleteAdvertAdmin(ctx context.Context, id int) error {
	return s.advertRepo.DeleteAdvert(ctx, id)
}
