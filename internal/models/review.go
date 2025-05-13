package models

import (
    "errors"
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Review struct {
    ID        int            `gorm:"primaryKey;autoIncrement"`
    UserID    uuid.UUID      `gorm:"type:uuid;not null;index"` // Tambahkan indeks pada UserID
    User      User           `gorm:"foreignKey:UserID"`
    EventID   int            `gorm:"not null;index"`          // Tambahkan indeks pada EventID
    Event     Event          `gorm:"foreignKey:EventID"`
    Rating    int            `gorm:"not null;check:rating >= 1 AND rating <= 5"` // Validasi rentang nilai
    Comment   string         `gorm:"type:text;not null"`                         // Pastikan komentar tidak kosong
    CreatedAt time.Time      `gorm:"autoCreateTime"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime"` // Metadata tambahan untuk pembaruan
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

// BeforeSave adalah hook untuk memvalidasi nilai Rating sebelum menyimpan ulasan
func (r *Review) BeforeSave(tx *gorm.DB) (err error) {
    if r.Rating < 1 || r.Rating > 5 {
        return errors.New("rating harus berada dalam rentang 1 hingga 5")
    }
    return nil
}