package repository

import (
    "ticketing-konser/internal/models"
    "gorm.io/gorm"
)

type NotificationRepository struct {
    db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
    return &NotificationRepository{db: db}
}

func (r *NotificationRepository) Create(notification *models.Notification) error {
    return r.db.Create(notification).Error
}

func (r *NotificationRepository) FindByUserID(userID string) ([]models.Notification, error) {
    var notifications []models.Notification
    if err := r.db.Where("user_id = ?", userID).Find(&notifications).Error; err != nil {
        return nil, err
    }
    return notifications, nil
}