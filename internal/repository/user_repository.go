package repository

import (
    "ticketing-konser/internal/models"
    "gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

// FindByEmail mencari user berdasarkan email
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
    var user models.User
    if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

// Create menyimpan user baru ke database
func (r *UserRepository) Create(user *models.User) error {
    return r.db.Create(user).Error
}

// FindRoleByName mencari role berdasarkan nama
func (r *UserRepository) FindRoleByName(name string) (*models.Role, error) {
    var role models.Role
    if err := r.db.Where("name = ?", name).First(&role).Error; err != nil {
        return nil, err
    }
    return &role, nil
}

// FindByID mencari user berdasarkan ID
func (r *UserRepository) FindByID(id string) (*models.User, error) {
    var user models.User
    if err := r.db.First(&user, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

// Update memperbarui data user
func (r *UserRepository) Update(user *models.User) error {
    return r.db.Save(user).Error
}

// Delete menghapus user berdasarkan ID
func (r *UserRepository) Delete(id string) error {
    return r.db.Delete(&models.User{}, "id = ?", id).Error
}