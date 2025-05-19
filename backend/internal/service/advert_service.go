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

func (s *AdvertService) GetAllAdverts(ctx context.Context, filter *dto.AdvertFilter) ([]*dto.GetAdvertResponse, int64, error) {
	return s.repo.GetAllAdverts(ctx, filter)
}
func (s *AdvertService) GetAdvertById(ctx context.Context, id int) (*dto.GetAdvertResponse, error) {
	return s.repo.GetAdvertById(ctx, id)
}
func (s *AdvertService) GetAllUserAdverts(ctx context.Context, userId int) ([]*dto.GetAdvertResponse, error) {
	return s.repo.GetAllUserAdverts(ctx, userId)
}
func (s *AdvertService) GetUserAdvertById(ctx context.Context, userId int, id int) (*dto.GetAdvertResponse, error) {
	return s.repo.GetUserAdvertById(ctx, userId, id)
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
func (s *AdvertService) GetAllAdvertsAdmin(ctx context.Context) ([]*dto.GetAdvertResponse, error) {
	return s.repo.GetAllAdvertsAdmin(ctx)
}

func (s *AdvertService) GetAdvertByIdAdmin(ctx context.Context, id int) (*dto.GetAdvertResponse, error) {
	return s.repo.GetAdvertByIdAdmin(ctx, id)
}

func (s *AdvertService) UpdateAdvertAdmin(ctx context.Context, id int, input *dto.UpdateAdvertInput) error {
	return s.repo.UpdateAdvertAdmin(ctx, id, input)
}

func (s *AdvertService) DeleteAdvertAdmin(ctx context.Context, id int) error {
	return s.repo.DeleteAdvertAdmin(ctx, id)
}
