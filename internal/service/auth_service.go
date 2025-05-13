package service

import (
	"errors"
	"log"
	"ticketing-konser/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

// Konstanta untuk pesan error
var (
	ErrInvalidCredentials  = errors.New("email atau password salah")
	ErrTokenCreation       = errors.New("gagal membuat token")
	ErrInvalidRefreshToken = errors.New("refresh token tidak valid")
)

// AuthService menggunakan interface agar mudah dimock saat testing
type AuthService struct {
	userService IUserService
	jwtUtil     utils.IJWTUtil
}

// NewAuthService menerima interface, bukan tipe konkret
func NewAuthService(userService IUserService, jwtUtil utils.IJWTUtil) *AuthService {
	return &AuthService{
		userService: userService,
		jwtUtil:     jwtUtil,
	}
}

// Login melakukan autentikasi user dan mengembalikan token JWT
func (s *AuthService) Login(email, password string) (string, error) {
	if email == "" || password == "" {
		return "", ErrInvalidCredentials
	}

	user, err := s.userService.GetUserByEmail(email)
	if err != nil {
		log.Printf("Login gagal: user dengan email %s tidak ditemukan", email)
		return "", ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Printf("Login gagal: password tidak cocok untuk email %s", email)
		return "", ErrInvalidCredentials
	}

	token, err := s.jwtUtil.GenerateToken(user.ID.String(), user.Role.Name)
	if err != nil {
		log.Printf("Login gagal: gagal membuat token untuk email %s", email)
		return "", ErrTokenCreation
	}

	log.Printf("Login berhasil untuk email %s", email)
	return token, nil
}

// RefreshToken membuat token baru berdasarkan refresh token yang valid
func (s *AuthService) RefreshToken(refreshToken string) (string, error) {
	claims, err := s.jwtUtil.ValidateRefreshToken(refreshToken)
	if err != nil {
		log.Printf("Refresh token gagal: %v", err)
		return "", ErrInvalidRefreshToken
	}

	token, err := s.jwtUtil.GenerateToken(claims.UserID, claims.Role)
	if err != nil {
		log.Printf("Refresh token gagal: gagal membuat token baru untuk user %s", claims.UserID)
		return "", ErrTokenCreation
	}

	log.Printf("Refresh token berhasil untuk user %s", claims.UserID)
	return token, nil
}
