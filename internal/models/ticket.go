package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Ticket adalah model untuk tiket
type Ticket struct {
	ID           int            `gorm:"primaryKey;autoIncrement"`
	UserID       uuid.UUID      `gorm:"type:uuid;not null;index"` // Tambahkan indeks pada UserID
	User         User           `gorm:"foreignKey:UserID"`
	EventID      int            `gorm:"not null;index"` // Tambahkan indeks pada EventID
	Event        Event          `gorm:"foreignKey:EventID"`
	Status       Status         `gorm:"size:50;not null;default:'active'"`           // Gunakan enum untuk Status
	TicketType   string         `gorm:"size:50;not null"`                            // Metadata tambahan
	Price        float64        `gorm:"type:decimal(15,2);not null;check:price > 0"` // Validasi nilai positif
	PurchaseDate time.Time      `gorm:"not null"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

// BeforeSave adalah hook untuk memvalidasi nilai Status sebelum menyimpan tiket
func (t *Ticket) BeforeSave(tx *gorm.DB) (err error) {
	if !IsValidStatus(t.Status) { // Pakai fungsi dari status.go
		return errors.New("status tiket tidak valid")
	}
	return nil
}
