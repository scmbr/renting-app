package dto

import (
	"time"

	"github.com/scmbr/renting-app/internal/models"
)

type CreateReviewInput struct {
	TargetID uint   `json:"target_id" binding:"required"`
	Rating   int    `json:"rating" binding:"required,min=1,max=5"`
	Comment  string `json:"comment"`
}

type GetReviewResponse struct {
	ID        uint           `json:"id"`
	AuthorID  uint           `json:"author_id"`
	TargetID  uint           `json:"target_id"`
	Rating    int            `json:"rating"`
	Comment   string         `json:"comment"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Author    *GetUserPublic `json:"author"`
}

type UpdateReviewInput struct {
	Rating  *int   `json:"rating" binding:"omitempty,min=1,max=5"`
	Comment *string `json:"comment"`
}

type GetUserPublic struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Avatar   string `json:"profile_picture"`
	Rating   float32 `json:"rating"`
}

func FromReview(r models.Review) *GetReviewResponse {
	return &GetReviewResponse{
		ID:        r.ID,
		AuthorID:  r.AuthorID,
		TargetID:  r.TargetID,
		Rating:    r.Rating,
		Comment:   r.Comment,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
		Author: &GetUserPublic{
			ID:       r.Author.ID,
			Name:     r.Author.Name,
			Surname:  r.Author.Surname,
			Avatar:   r.Author.ProfilePicture,
			Rating:   r.Author.Rating,
		},
	}
}
