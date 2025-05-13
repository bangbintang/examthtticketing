package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TransactionStatus adalah alias dari Status
type TransactionStatus Status

// Transaction adalah model untuk transaksi
type Transaction struct {
	ID              int               `gorm:"primaryKey;autoIncrement"`
	UserID          uuid.UUID         `gorm:"type:uuid;not null"`
	User            User              `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	EventID         int               `gorm:"not null"`
	Event           Event             `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	TicketID        int               `gorm:"not null"`
	Ticket          Ticket            `gorm:"foreignKey:TicketID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Amount          float64           `gorm:"type:decimal(15,2);not null;check:amount > 0"`
	Status          TransactionStatus `gorm:"size:50;not null;default:'pending'"`
	TransactionDate time.Time         `gorm:"not null"`
	CreatedAt       time.Time         `gorm:"autoCreateTime"`
	UpdatedAt       time.Time         `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt    `gorm:"index"`
}

// BeforeSave adalah hook untuk memvalidasi nilai sebelum menyimpan transaksi
func (t *Transaction) BeforeSave(tx *gorm.DB) (err error) {
	// Validasi Status
	if !isValidTransactionStatus(t.Status) {
		return errors.New("status transaksi tidak valid")
	}

	// Validasi Amount
	if t.Amount <= 0 {
		return errors.New("amount harus lebih besar dari 0")
	}

	// Validasi TransactionDate
	if t.TransactionDate.IsZero() {
		return errors.New("transaction date tidak boleh kosong")
	}

	// Validasi UserID
	if t.UserID == uuid.Nil {
		return errors.New("user ID tidak valid")
	}

	return nil
}

// isValidTransactionStatus memeriksa apakah nilai status valid
func isValidTransactionStatus(status TransactionStatus) bool {
	validStatuses := []TransactionStatus{
		TransactionStatus(StatusPending),
		TransactionStatus(StatusCompleted),
		TransactionStatus(StatusFailed),
	}
	for _, validStatus := range validStatuses {
		if status == validStatus {
			return true
		}
	}
	return false
}
