package service

import (
	"errors"
	"ticketing-konser/internal/models"
	"ticketing-konser/internal/repository"
)

type NotificationService struct {
	notificationRepo *repository.NotificationRepository
}

func NewNotificationService(notificationRepo *repository.NotificationRepository) *NotificationService {
	return &NotificationService{notificationRepo: notificationRepo}
}

func (s *NotificationService) CreateNotification(notification *models.Notification) (*models.Notification, error) {
	if notification.Message == "" {
		return nil, errors.New("pesan notifikasi tidak boleh kosong")
	}
	if err := s.notificationRepo.Create(notification); err != nil {
		return nil, err
	}
	return notification, nil
}

func (s *NotificationService) GetNotificationsByUser(userID string) ([]models.Notification, error) {
	return s.notificationRepo.FindByUserID(userID)
}
