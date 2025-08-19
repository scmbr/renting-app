package repository

import (
	"context"

	"github.com/scmbr/renting-app/internal/domain"
	"github.com/scmbr/renting-app/internal/dto"
	"gorm.io/gorm"
)

type ReviewRepo struct {
	db *gorm.DB
}

func NewReviewRepo(db *gorm.DB) *ReviewRepo {
	return &ReviewRepo{db: db}
}

func (r *ReviewRepo) Create(ctx context.Context, authorID uint, input dto.CreateReviewInput) (*dto.GetReviewResponse, error) {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	review := domain.Review{
		AuthorID: authorID,
		TargetID: input.TargetID,
		Rating:   input.Rating,
		Comment:  input.Comment,
	}

	if err := tx.Create(&review).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Preload("Author").First(&review, review.ID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return dto.FromReview(review), nil
}

func (r *ReviewRepo) Update(ctx context.Context, reviewID uint, input dto.UpdateReviewInput) (*dto.GetReviewResponse, error) {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var review domain.Review
	if err := tx.First(&review, reviewID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if input.Rating != nil {
		review.Rating = *input.Rating
	}
	if input.Comment != nil {
		review.Comment = *input.Comment
	}

	if err := tx.Save(&review).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Preload("Author").First(&review, review.ID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return dto.FromReview(review), nil
}

func (r *ReviewRepo) Delete(ctx context.Context, reviewID uint) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Delete(&domain.Review{}, reviewID).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *ReviewRepo) GetByAuthorID(ctx context.Context, authorID uint) ([]*dto.GetReviewResponse, error) {
	var reviews []domain.Review

	err := r.db.WithContext(ctx).
		Preload("Author").
		Where("author_id = ?", authorID).
		Find(&reviews).Error

	if err != nil {
		return nil, err
	}

	var result []*dto.GetReviewResponse
	for _, review := range reviews {
		result = append(result, dto.FromReview(review))
	}

	return result, nil
}

func (r *ReviewRepo) GetByTargetID(ctx context.Context, targetID uint) ([]*dto.GetReviewResponse, error) {
	var reviews []domain.Review

	err := r.db.WithContext(ctx).
		Preload("Author").
		Where("target_id = ?", targetID).
		Find(&reviews).Error

	if err != nil {
		return nil, err
	}

	var result []*dto.GetReviewResponse
	for _, review := range reviews {
		result = append(result, dto.FromReview(review))
	}

	return result, nil
}

func (r *ReviewRepo) ReviewExists(ctx context.Context, authorID uint, targetID uint) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&domain.Review{}).
		Where("author_id = ? AND target_id = ?", authorID, targetID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *ReviewRepo) IsAuthor(ctx context.Context, userID uint, reviewID uint) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&domain.Review{}).
		Where("id = ? AND author_id = ?", reviewID, userID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
