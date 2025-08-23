package service

import (
	"context"
	"mime/multipart"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/repository"
	"github.com/scmbr/renting-app/pkg/storage"
	"github.com/sirupsen/logrus"
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
	}, "apartments_photo")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"filename": fileHeader.Filename,
			"size":     fileHeader.Size,
			"type":     fileHeader.Header.Get("Content-Type"),
		}).Errorf("failed to upload to S3: %s", err.Error())
		return "", err
	}

	return url, nil
}
func (s *ApartmentPhotoService) GetAllPhotos(ctx context.Context, apartmentId int) ([]dto.GetApartmentPhotoResponse, error) {
	return s.repo.GetAllPhotos(ctx, apartmentId)
}

func (s *ApartmentPhotoService) GetPhotoById(ctx context.Context, apartmentId, photoId int) (*dto.GetApartmentPhotoResponse, error) {
	return s.repo.GetPhotoById(ctx, apartmentId, photoId)
}

func (s *ApartmentPhotoService) AddPhotos(ctx context.Context, userId, apartmentId int, files []*multipart.FileHeader) ([]*dto.GetApartmentPhotoResponse, error) {
	var inputs []dto.CreatePhotoInput

	hasCover, err := s.repo.HasCoverPhoto(ctx, apartmentId)
	if err != nil {
		return nil, err
	}

	for i, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		url, err := s.storage.Upload(ctx, storage.UploadInput{
			File:        file,
			Name:        fileHeader.Filename,
			Size:        fileHeader.Size,
			ContentType: fileHeader.Header.Get("Content-Type"),
		}, "apartments_photo")
		if err != nil {
			return nil, err
		}

		inputs = append(inputs, dto.CreatePhotoInput{
			ApartmentID: uint(apartmentId),
			URL:         url,
			FileName:    fileHeader.Filename,
			IsCover:     i == 0 && !hasCover,
		})
	}

	photos, err := s.repo.AddPhotos(ctx, userId, apartmentId, inputs)
	if err != nil {
		return nil, err
	}

	resp := make([]*dto.GetApartmentPhotoResponse, len(photos))
	for i, photo := range photos {
		result := dto.FromApartmentPhoto(photo)
		resp[i] = &result
	}

	return resp, nil
}

func (s *ApartmentPhotoService) DeletePhoto(ctx context.Context, userId, apartmentId, photoId int) error {
	return s.repo.DeletePhoto(ctx, userId, apartmentId, photoId)
}

func (s *ApartmentPhotoService) SetCover(ctx context.Context, userId, apartmentId, photoId int) error {
	return s.repo.SetCover(ctx, userId, apartmentId, photoId)
}
func (s *ApartmentPhotoService) HasCoverPhoto(ctx context.Context, apartmentId int) (bool, error) {
	return s.repo.HasCoverPhoto(ctx, apartmentId)
}
func (s *ApartmentPhotoService) ReplaceAllPhotos(ctx context.Context, userId, apartmentId int, inputs []dto.CreatePhotoInput) error {
	return s.repo.ReplaceAllPhotos(ctx, userId, apartmentId, inputs)
}
