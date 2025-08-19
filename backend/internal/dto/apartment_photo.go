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
type GetApartmentPhoto struct {
	ID          uint   `json:"id"`
	ApartmentID uint   `json:"apartment_id"`
	URL         string `json:"url"`
	IsCover     bool   `json:"is_cover"`
}

func FromApartmentPhoto(apartmentPhoto *domain.ApartmentPhoto) *GetApartmentPhoto {
	if apartmentPhoto == nil {
		return nil
	}

	return &GetApartmentPhoto{
		ID:          apartmentPhoto.ID,
		ApartmentID: apartmentPhoto.ApartmentID,
		URL:         apartmentPhoto.URL,
		IsCover:     apartmentPhoto.IsCover,
	}

}
