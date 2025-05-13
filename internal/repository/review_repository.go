package repository

import (
	"ticketing-konser/internal/models"

	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) Create(review *models.Review) error {
	return r.db.Create(review).Error
}

func (r *ReviewRepository) FindByEventID(eventID int) ([]models.Review, error) {
	var reviews []models.Review
	if err := r.db.Where("event_id = ?", eventID).Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *ReviewRepository) FindByID(id int) (*models.Review, error) {
	var review models.Review
	if err := r.db.First(&review, id).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *ReviewRepository) Update(review *models.Review) error {
	return r.db.Save(review).Error
}

func (r *ReviewRepository) Delete(id int) error {
	return r.db.Delete(&models.Review{}, id).Error
}
