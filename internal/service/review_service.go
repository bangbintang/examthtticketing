package service

import (
	"errors"
	"ticketing-konser/internal/models"
	"ticketing-konser/internal/repository"
)

// (Opsional) Interface agar mudah di-mock dan digunakan di handler/service lain
type IReviewService interface {
	CreateReview(review *models.Review) (*models.Review, error)
	GetReviewsByEvent(eventID int) ([]models.Review, error)
	GetReviewByID(id int) (*models.Review, error)
	UpdateReview(review *models.Review) error
	DeleteReview(id int) error
}

type ReviewService struct {
	reviewRepo *repository.ReviewRepository
}

func NewReviewService(reviewRepo *repository.ReviewRepository) *ReviewService {
	return &ReviewService{reviewRepo: reviewRepo}
}

func (s *ReviewService) CreateReview(review *models.Review) (*models.Review, error) {
	if review.Rating < 1 || review.Rating > 5 {
		return nil, errors.New("rating harus antara 1 sampai 5")
	}
	if err := s.reviewRepo.Create(review); err != nil {
		return nil, err
	}
	return review, nil
}

func (s *ReviewService) GetReviewsByEvent(eventID int) ([]models.Review, error) {
	return s.reviewRepo.FindByEventID(eventID)
}

func (s *ReviewService) GetReviewByID(id int) (*models.Review, error) {
	return s.reviewRepo.FindByID(id)
}

func (s *ReviewService) UpdateReview(review *models.Review) error {
	return s.reviewRepo.Update(review)
}

func (s *ReviewService) DeleteReview(id int) error {
	return s.reviewRepo.Delete(id)
}
