package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// NotificationType adalah tipe data khusus untuk jenis notifikasi
type NotificationType string

// Nilai-nilai yang valid untuk NotificationType
const (
	NotificationTypeInfo    NotificationType = "info"
	NotificationTypeWarning NotificationType = "warning"
	NotificationTypeError   NotificationType = "error"
)

type Notification struct {
	ID        int              `gorm:"primaryKey;autoIncrement"`
	UserID    uuid.UUID        `gorm:"type:uuid;not null;index"` // Tambahkan indeks pada UserID
	User      User             `gorm:"foreignKey:UserID"`
	Message   string           `gorm:"type:text;not null"`              // Validasi untuk memastikan Message tidak kosong
	Type      NotificationType `gorm:"size:50;not null;default:'info'"` // Enum untuk jenis notifikasi
	IsRead    bool             `gorm:"not null;default:false"`
	CreatedAt time.Time        `gorm:"autoCreateTime"`
	UpdatedAt time.Time        `gorm:"autoUpdateTime"` // Metadata tambahan untuk pembaruan
	DeletedAt gorm.DeletedAt   `gorm:"index"`
}

// BeforeSave adalah hook untuk memvalidasi nilai Type sebelum menyimpan notifikasi
func (n *Notification) BeforeSave(tx *gorm.DB) (err error) {
	if !isValidNotificationType(n.Type) {
		return errors.New("jenis notifikasi tidak valid")
	}
	return nil
}

// isValidNotificationType memeriksa apakah nilai Type valid
func isValidNotificationType(notificationType NotificationType) bool {
	switch notificationType {
	case NotificationTypeInfo, NotificationTypeWarning, NotificationTypeError:
		return true
	default:
		return false
	}
}
