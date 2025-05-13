package utils

import (
    "github.com/stretchr/testify/mock"
)

type MockJWTUtil struct {
    mock.Mock
}

// GenerateToken adalah mock untuk metode GenerateToken
func (m *MockJWTUtil) GenerateToken(userID, role string) (string, error) {
    args := m.Called(userID, role)
    return args.String(0), args.Error(1)
}

// ValidateToken adalah mock untuk metode ValidateToken
func (m *MockJWTUtil) ValidateToken(token string) (*CustomClaims, error) {
    args := m.Called(token)
    claims := args.Get(0)
    if claims == nil {
        return nil, args.Error(1)
    }
    return claims.(*CustomClaims), args.Error(1)
}

// GenerateRefreshToken adalah mock untuk metode GenerateRefreshToken
func (m *MockJWTUtil) GenerateRefreshToken(userID, role string) (string, error) {
    args := m.Called(userID, role)
    return args.String(0), args.Error(1)
}

// ValidateRefreshToken adalah mock untuk metode ValidateRefreshToken
func (m *MockJWTUtil) ValidateRefreshToken(token string) (*CustomClaims, error) {
    args := m.Called(token)
    claims := args.Get(0)
    if claims == nil {
        return nil, args.Error(1)
    }
    return claims.(*CustomClaims), args.Error(1)
}