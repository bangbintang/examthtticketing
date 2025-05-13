package service

import (
    "errors"
    "ticketing-konser/internal/models"
    "ticketing-konser/internal/repository"
)

type EventService struct {
    eventRepo *repository.EventRepository
}

func NewEventService(eventRepo *repository.EventRepository) *EventService {
    return &EventService{eventRepo: eventRepo}
}

func (s *EventService) CreateEvent(event *models.Event) (*models.Event, error) {
    if event.Name == "" {
        return nil, errors.New("nama event tidak boleh kosong")
    }
    if err := s.eventRepo.Create(event); err != nil {
        return nil, err
    }
    return event, nil
}

func (s *EventService) GetAllEvents() ([]models.Event, error) {
    return s.eventRepo.FindAll()
}

func (s *EventService) GetEventByID(id int) (*models.Event, error) {
    return s.eventRepo.FindByID(id)
}

func (s *EventService) UpdateEvent(event *models.Event) error {
    return s.eventRepo.Update(event)
}

func (s *EventService) DeleteEvent(id int) error {
    return s.eventRepo.Delete(id)
}