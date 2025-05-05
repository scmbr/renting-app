package service

import (
	"context"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/repository"
)

type AdvertService struct {
	repo repository.Advert
}

func NewAdvertService(repo repository.Advert) *AdvertService {
	return &AdvertService{
		repo: repo,
	}
}
func (s *AdvertService) GetAllAdverts(ctx context.Context, userId int) ([]*dto.GetAdvertResponse, error) {
	return s.repo.GetAllAdverts(ctx, userId)
}
func (s *AdvertService) GetAdvertById(ctx context.Context, userId int, id int) (*dto.GetAdvertResponse, error) {
	return s.repo.GetAdvertById(ctx, userId, id)
}
func (s *AdvertService) CreateAdvert(ctx context.Context, userId int, input dto.CreateAdvertInput) error {
	return s.repo.CreateAdvert(ctx, userId, input)
}
func (s *AdvertService) DeleteAdvert(ctx context.Context, userId int, id int) error {
	return s.repo.DeleteAdvert(ctx, userId, id)
}
func (s *AdvertService) UpdateAdvert(ctx context.Context, userId int, id int, input *dto.UpdateAdvertInput) error {
	return s.repo.UpdateAdvert(ctx, userId, id, input)
}
