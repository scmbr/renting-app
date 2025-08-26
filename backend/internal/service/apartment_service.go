package service

import (
	"context"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/repository"
)

type ApartmentService struct {
	repo repository.Apartment
	apartmentPhotoRepo repository.ApartmentPhoto
}

func NewApartmentService(repo repository.Apartment, apartmentPhotoRepo repository.ApartmentPhoto) *ApartmentService {
	return &ApartmentService{
		repo: repo,
		apartmentPhotoRepo: apartmentPhotoRepo,
	}
}

func (s *ApartmentService) GetAllApartments(ctx context.Context, userId int) ([]*dto.GetApartmentResponse, error) {
	apartments, err:= s.repo.GetAllApartments(ctx, userId)
	if err!=nil{
		return nil,err
	}
	apartmentIDs := make([]uint, len(apartments))
	for i, aps := range apartments {
		apartmentIDs[i] = aps.ID
	}
	result := make([]*dto.GetApartmentResponse, len(apartments))
	photosMap, err := s.apartmentPhotoRepo.GetAllPhotosForApartments(ctx, apartmentIDs)
	if err != nil {
		return nil,  err
	}
	for i, aps := range apartments {
		resp := dto.FromApartment(aps)

		resp.ApartmentPhotos = photosMap[aps.ID]
		if resp.ApartmentPhotos == nil {
			resp.ApartmentPhotos = []dto.GetApartmentPhotoResponse{}
		}
		result[i] = resp
	}
	return result,nil
}
func (s *ApartmentService) GetApartmentById(ctx context.Context, userId int, id int) (*dto.GetApartmentResponse, error) {
	return s.repo.GetApartmentById(ctx, userId, id)
}
func (s *ApartmentService) CreateApartment(ctx context.Context, userId int, input dto.CreateApartmentInput) (*dto.GetApartmentResponse, error) {
	apartment, err := s.repo.CreateApartment(ctx, userId, input)
	if err != nil {
		return nil, err
	}
	resp := dto.FromApartment(apartment)
	return resp, nil
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
