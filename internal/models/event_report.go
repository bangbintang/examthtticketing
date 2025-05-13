package models

import (
    "time"

    "gorm.io/gorm"
)

type EventReport struct {
    ID          int            `gorm:"primaryKey;autoIncrement"`
    EventID     int            `gorm:"not null;index"` // Tambahkan indeks pada EventID
    Event       Event          `gorm:"foreignKey:EventID"`
    TicketsSold int            `gorm:"not null;default:0;check:tickets_sold >= 0"` // Validasi nilai positif
    Revenue     float64        `gorm:"type:decimal(15,2);not null;default:0;check:revenue >= 0"` // Validasi nilai positif
    ReportDate  time.Time      `gorm:"type:date;not null;index"` // Tambahkan indeks pada ReportDate
    CreatedAt   time.Time      `gorm:"autoCreateTime"`
    UpdatedAt   time.Time      `gorm:"autoUpdateTime"` // Metadata tambahan untuk pembaruan
    DeletedAt   gorm.DeletedAt `gorm:"index"`
}