package service

import (
	"context"
	"fmt"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/repository"
)

type ApartmentService struct {
	repo repository.Apartment
}

func NewApartmentService(repo repository.Apartment) *ApartmentService {
	return &ApartmentService{
		repo: repo,
	}
}

func (s *ApartmentService) GetAllApartments(ctx context.Context, userId int) ([]*dto.GetApartmentResponse, error) {
	return s.repo.GetAllApartments(ctx, userId)
}
func (s *ApartmentService) GetApartmentById(ctx context.Context, userId int, id int) (*dto.GetApartmentResponse, error) {
	return s.repo.GetApartmentById(ctx, userId, id)
}
func (s *ApartmentService) CreateApartment(ctx context.Context, userId int, input dto.CreateApartmentInput) error {
	fmt.Println(input.City)
	return s.repo.CreateApartment(ctx, userId, input)
}
func (s *ApartmentService) DeleteApartment(ctx context.Context, userId int, id int) error {
	return s.repo.DeleteApartment(ctx, userId, id)
}
func (s *ApartmentService) UpdateApartment(ctx context.Context, userId int, id int, input *dto.UpdateApartmentInput) error {
	return s.repo.UpdateApartment(ctx, userId, id, input)
}
func (s *ApartmentService) GetAllApartmentsAdmin(ctx context.Context) ([]*dto.GetApartmentResponse, error) {
	return s.repo.GetAllApartmentsAdmin(ctx)
}

func (s *ApartmentService) GetApartmentByIdAdmin(ctx context.Context, id int) (*dto.GetApartmentResponse, error) {
	return s.repo.GetApartmentByIdAdmin(ctx, id)
}

func (s *ApartmentService) UpdateApartmentAdmin(ctx context.Context, id int, input *dto.UpdateApartmentInput) error {
	return s.repo.UpdateApartmentAdmin(ctx, id, input)
}

func (s *ApartmentService) DeleteApartmentAdmin(ctx context.Context, id int) error {
	return s.repo.DeleteApartmentAdmin(ctx, id)
}
