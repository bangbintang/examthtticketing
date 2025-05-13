package repository

import (
    "ticketing-konser/internal/models"
    "gorm.io/gorm"
)

type EventRepository struct {
    db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
    return &EventRepository{db: db}
}

// Create menyimpan event baru ke database
func (r *EventRepository) Create(event *models.Event) error {
    return r.db.Create(event).Error
}

// FindAll mengambil semua event dari database
func (r *EventRepository) FindAll() ([]models.Event, error) {
    var events []models.Event
    if err := r.db.Find(&events).Error; err != nil {
        return nil, err
    }
    return events, nil
}

// FindByID mengambil event berdasarkan ID
func (r *EventRepository) FindByID(id int) (*models.Event, error) {
    var event models.Event
    if err := r.db.First(&event, id).Error; err != nil {
        return nil, err
    }
    return &event, nil
}

// Update memperbarui data event
func (r *EventRepository) Update(event *models.Event) error {
    return r.db.Save(event).Error
}

// Delete menghapus event berdasarkan ID
func (r *EventRepository) Delete(id int) error {
    return r.db.Delete(&models.Event{}, id).Error
}