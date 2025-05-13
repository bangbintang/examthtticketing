package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AuditAction adalah tipe data khusus untuk jenis tindakan audit
type AuditAction string

// Nilai-nilai yang valid untuk AuditAction
const (
	ActionCreate AuditAction = "create"
	ActionUpdate AuditAction = "update"
	ActionDelete AuditAction = "delete"
	ActionLogin  AuditAction = "login"
	ActionLogout AuditAction = "logout"
)

type AuditTrail struct {
	ID          int            `gorm:"primaryKey;autoIncrement"`
	UserID      uuid.UUID      `gorm:"type:uuid;not null;index"` // Tambahkan indeks pada UserID
	User        User           `gorm:"foreignKey:UserID"`
	Action      AuditAction    `gorm:"size:50;not null"` // Gunakan enum untuk Action
	Description string         `gorm:"type:text"`
	IPAddress   string         `gorm:"size:45;not null"` // Metadata tambahan untuk IP Address
	Timestamp   time.Time      `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// BeforeSave adalah hook untuk memvalidasi nilai Action sebelum menyimpan audit trail
func (a *AuditTrail) BeforeSave(tx *gorm.DB) (err error) {
	if !isValidAuditAction(a.Action) {
		return errors.New("jenis tindakan audit tidak valid")
	}
	return nil
}

// isValidAuditAction memeriksa apakah nilai Action valid
func isValidAuditAction(action AuditAction) bool {
	switch action {
	case ActionCreate, ActionUpdate, ActionDelete, ActionLogin, ActionLogout:
		return true
	default:
		return false
	}
}
