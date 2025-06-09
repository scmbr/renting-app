package service

import (
	"context"
	"mime/multipart"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/repository"
	"github.com/scmbr/renting-app/pkg/storage"
)

type ApartmentPhotoService struct {
	repo    repository.ApartmentPhoto
	storage storage.Provider
}

func NewApartmentPhotoService(repo repository.ApartmentPhoto, storage storage.Provider) *ApartmentPhotoService {
	return &ApartmentPhotoService{
		repo:    repo,
		storage: storage,
	}
}
func (s *ApartmentPhotoService) UploadPhotoToS3(ctx context.Context, fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()
	url, err := s.storage.Upload(ctx, storage.UploadInput{
		File:        file,
		Name:        fileHeader.Filename,
		Size:        fileHeader.Size,
		ContentType: fileHeader.Header.Get("Content-Type"),
	})
	if err != nil {
		return "", err
	}

	return url, nil
}
func (s *ApartmentPhotoService) GetAllPhotos(ctx context.Context, apartmentId int) ([]*dto.GetApartmentPhoto, error) {
	return s.repo.GetAllPhotos(ctx, apartmentId)
}

func (s *ApartmentPhotoService) GetPhotoById(ctx context.Context, apartmentId, photoId int) (*dto.GetApartmentPhoto, error) {
	return s.repo.GetPhotoById(ctx, apartmentId, photoId)
}

func (s *ApartmentPhotoService) AddPhotoBatch(ctx context.Context, userId, apartmentId int, inputs []dto.CreatePhotoInput) error {
	return s.repo.AddPhotoBatch(ctx, userId, apartmentId, inputs)
}

func (s *ApartmentPhotoService) DeletePhoto(ctx context.Context, userId, apartmentId, photoId int) error {
	return s.repo.DeletePhoto(ctx, userId, apartmentId, photoId)
}

func (s *ApartmentPhotoService) SetCover(ctx context.Context, userId, apartmentId, photoId int) error {
	return s.repo.SetCover(ctx, userId, apartmentId, photoId)
}
func (s *ApartmentPhotoService) HasCoverPhoto(apartmentId int) (bool, error) {
	return s.repo.HasCoverPhoto(apartmentId)
}
func (s *ApartmentPhotoService) ReplaceAllPhotos(ctx context.Context, userId, apartmentId int, inputs []dto.CreatePhotoInput) error {
	return s.repo.ReplaceAllPhotos(ctx, userId, apartmentId, inputs)
}