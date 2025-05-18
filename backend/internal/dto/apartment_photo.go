package dto

import "mime/multipart"

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
