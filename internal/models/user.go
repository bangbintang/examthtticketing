package models

import (
	"errors"
	"time"

	"ticketing-konser/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey"` // UUID dihasilkan di aplikasi
	Name      string         `gorm:"size:255;not null"`
	Email     string         `gorm:"size:255;uniqueIndex;not null"`
	Password  string         `gorm:"size:255;not null"`
	RoleID    int            `gorm:"not null;index"` // Tambahkan indeks pada RoleID
	Role      Role           `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	LastLogin *time.Time     `gorm:"default:null"` // Metadata tambahan untuk login terakhir
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Role struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"size:100;not null;unique"`
}

// BeforeCreate adalah hook untuk menghasilkan UUID sebelum membuat user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate UUID jika ID tidak ada
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	// Validasi data sebelum disimpan
	if u.Name == "" {
		return errors.New("name tidak boleh kosong")
	}

	if !utils.IsValidEmail(u.Email) {
		return errors.New("email tidak valid")
	}

	if err := utils.ValidatePassword(u.Password); err != nil {
		return err
	}

	return nil
}
