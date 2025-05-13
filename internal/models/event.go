package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID          int            `gorm:"primaryKey;autoIncrement"`
	Name        string         `gorm:"size:255;not null;uniqueIndex"` // Indeks unik pada Name
	Description string         `gorm:"type:text"`
	Capacity    int            `gorm:"not null;check:capacity > 0"`                  // Validasi nilai positif
	Price       float64        `gorm:"type:decimal(10,2);not null;check:price >= 0"` // Validasi nilai positif
	Status      Status         `gorm:"size:50;not null;default:'active'"`            // Gunakan enum untuk Status
	Location    string         `gorm:"size:255;not null"`                            // Metadata tambahan untuk lokasi event
	StartDate   time.Time      `gorm:"type:date;not null"`
	EndDate     time.Time      `gorm:"type:date;not null"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// BeforeSave adalah hook untuk memvalidasi nilai Status sebelum menyimpan event
func (e *Event) BeforeSave(tx *gorm.DB) (err error) {
	if !IsValidStatus(e.Status) { // Pakai fungsi dari status.go
		return errors.New("status event tidak valid")
	}
	return nil
}

// Hapus fungsi isValidEventStatus, gunakan IsValidStatus dari status.go
