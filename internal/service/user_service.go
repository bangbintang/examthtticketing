// filepath: [user_service.go](http://_vscodecontentref_/4)
package service

import (
	"errors"
	"ticketing-konser/internal/models"
	"ticketing-konser/internal/repository"
	"ticketing-konser/internal/utils"
)

// IUserService interface agar mudah di-mock dan digunakan di AuthService
type IUserService interface {
	RegisterUser(user *models.User) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) RegisterUser(user *models.User) (*models.User, error) {
	if !utils.IsValidEmail(user.Email) {
		return nil, errors.New("email tidak valid")
	}

	if err := utils.ValidatePassword(user.Password); err != nil {
		return nil, err
	}

	existingUser, err := s.userRepo.FindByEmail(user.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email sudah terdaftar")
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.FindByEmail(email)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.userRepo.Delete(id)
}
