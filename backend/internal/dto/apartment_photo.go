package dto

import (
	"mime/multipart"

	"github.com/scmbr/renting-app/internal/domain"
)

type CreatePhotoInput struct {
	ApartmentID uint                  `json:"apartment_id" binding:"required"`
	IsCover     bool                  `json:"is_cover"`
	Photo       *multipart.FileHeader `form:"photo" binding:"required"`
	URL         string
	FileName    string
}
type GetApartmentPhotoResponse struct {
	ID          uint   `json:"id"`
	ApartmentID uint   `json:"apartment_id"`
	URL         string `json:"url"`
	IsCover     bool   `json:"is_cover"`
}

func FromApartmentPhoto(apartmentPhoto *domain.ApartmentPhoto) GetApartmentPhotoResponse {
	if apartmentPhoto == nil {
		return GetApartmentPhotoResponse{}
	}

	return GetApartmentPhotoResponse{
		ID:          apartmentPhoto.ID,
		ApartmentID: apartmentPhoto.ApartmentID,
		URL:         apartmentPhoto.URL,
		IsCover:     apartmentPhoto.IsCover,
	}
}
