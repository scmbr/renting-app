package dto

import "github.com/scmbr/renting-app/internal/domain"

type GetProfileResponse struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	Surname        string  `json:"surname"`
	ProfilePicture string  `json:"profile_picture"`
	City           string  `json:"city,omitempty"`
	Rating         float32 `json:"rating"`
}

func FromUserToProfile(u *domain.User) *GetProfileResponse {
	return &GetProfileResponse{
		Id:             int(u.ID),
		Name:           u.Name,
		Surname:        u.Surname,
		ProfilePicture: u.ProfilePicture,
		City:           u.City,
		Rating:         u.Rating,
	}
}
