package service

import (
    "ticketing-konser/internal/models"

    "github.com/stretchr/testify/mock"
)

// MockUserService adalah mock untuk IUserService
// Digunakan untuk pengujian unit tanpa bergantung pada implementasi konkret
type MockUserService struct {
    mock.Mock
}

// GetUserByEmail adalah mock untuk metode GetUserByEmail
func (m *MockUserService) GetUserByEmail(email string) (*models.User, error) {
    args := m.Called(email)

    // Validasi tipe data untuk memastikan nilai yang dikembalikan adalah *models.User
    user, ok := args.Get(0).(*models.User)
    if !ok && args.Get(0) != nil {
        return nil, args.Error(1)
    }

    return user, args.Error(1)
}

// Contoh metode tambahan jika IUserService memiliki metode lain
func (m *MockUserService) CreateUser(user *models.User) error {
    args := m.Called(user)
    return args.Error(0)
}