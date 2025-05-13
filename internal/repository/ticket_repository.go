package repository

import (
    "ticketing-konser/internal/models"
    "gorm.io/gorm"
)

type TicketRepository struct {
    db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *TicketRepository {
    return &TicketRepository{db: db}
}

// Create menyimpan tiket baru ke database
func (r *TicketRepository) Create(ticket *models.Ticket) error {
    return r.db.Create(ticket).Error
}

// FindByUserID mengambil semua tiket berdasarkan ID pengguna
func (r *TicketRepository) FindByUserID(userID string) ([]models.Ticket, error) {
    var tickets []models.Ticket
    if err := r.db.Where("user_id = ?", userID).Find(&tickets).Error; err != nil {
        return nil, err
    }
    return tickets, nil
}

// FindByEventID mengambil semua tiket berdasarkan ID event
func (r *TicketRepository) FindByEventID(eventID int) ([]models.Ticket, error) {
    var tickets []models.Ticket
    if err := r.db.Where("event_id = ?", eventID).Find(&tickets).Error; err != nil {
        return nil, err
    }
    return tickets, nil
}

// Delete menghapus tiket berdasarkan ID
func (r *TicketRepository) Delete(ticketID int) error {
    return r.db.Delete(&models.Ticket{}, ticketID).Error
}