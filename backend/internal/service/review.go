package service

import (
	"context"
	"errors"
	"log"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/repository"
)

type ReviewService struct {
	reviewRepo repository.Review
	userRepo  repository.Users
}

func NewReviewService(reviewRepo repository.Review, userRepo repository.Users) *ReviewService {
	return &ReviewService{
		reviewRepo: reviewRepo,
		userRepo: userRepo,
	}
}

func (s *ReviewService) Create(ctx context.Context, authorID uint, input dto.CreateReviewInput) (*dto.GetReviewResponse, error) {
    if authorID == input.TargetID {
        return nil, errors.New("cannot review yourself")
    }


    exists, err := s.reviewRepo.ReviewExists(ctx, authorID, input.TargetID)
    if err != nil {
        return nil, err
    }
    if exists {
        return nil, errors.New("you have already reviewed this target")
    }

    review, err := s.reviewRepo.Create(ctx, authorID, input)
    if err != nil {
        return nil, err
    }

    err = s.updateUserRating(ctx, input.TargetID)
    if err != nil {
        log.Printf("failed to update user rating: %v", err)
    }

    return review, nil
}

func (s *ReviewService) updateUserRating(ctx context.Context, userID uint) error {
    reviews, err := s.reviewRepo.GetByTargetID(ctx, userID)
    if err != nil {
        return err
    }


    var sum float64
    for _, r := range reviews {
        sum += float64(r.Rating)
    }
    averageRating := sum / float64(len(reviews))


    err = s.userRepo.UpdateRating(ctx, userID, float32(averageRating))
    if err != nil {
        return err
    }

    return nil
}

func (s *ReviewService) Update(ctx context.Context, userID uint, reviewID uint, input dto.UpdateReviewInput) (*dto.GetReviewResponse, error) {
	isAuthor, err := s.reviewRepo.IsAuthor(ctx, userID, reviewID)
	if err != nil {
		return nil, err
	}
	if !isAuthor {
		return nil, errors.New("you are not the author of this review")
	}

	return s.reviewRepo.Update(ctx, reviewID, input)
}

func (s *ReviewService) Delete(ctx context.Context, userID uint, reviewID uint) error {
	isAuthor, err := s.reviewRepo.IsAuthor(ctx, userID, reviewID)
	if err != nil {
		return err
	}
	if !isAuthor {
		return errors.New("you are not the author of this review")
	}

	return s.reviewRepo.Delete(ctx, reviewID)
}

func (s *ReviewService) GetByAuthorID(ctx context.Context, authorID uint) ([]*dto.GetReviewResponse, error) {
	return s.reviewRepo.GetByAuthorID(ctx, authorID)
}

func (s *ReviewService) GetByTargetID(ctx context.Context, targetID uint) ([]*dto.GetReviewResponse, error) {
	return s.reviewRepo.GetByTargetID(ctx, targetID)
}