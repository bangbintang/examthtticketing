package service

import (
	"errors"
	"ticketing-konser/internal/models"
	"ticketing-konser/internal/repository"

	"github.com/google/uuid"
)

// ITransactionService agar mudah di-mock dan digunakan di service lain
type ITransactionService interface {
	CreateTransaction(transaction *models.Transaction) (*models.Transaction, error)
	GetTransactionsByUser(userID string) ([]models.Transaction, error)
	GetTransactionByID(transactionID int) (*models.Transaction, error)
	GetTransactionsByEvent(eventID int) ([]models.Transaction, error)
	UpdateTransaction(transaction *models.Transaction) error
	DeleteTransaction(transactionID int) error
}

type TransactionService struct {
	transactionRepo *repository.TransactionRepository
}

func NewTransactionService(transactionRepo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{transactionRepo: transactionRepo}
}

func (s *TransactionService) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	if transaction.UserID == uuid.Nil {
		return nil, errors.New("user ID tidak boleh kosong")
	}
	if transaction.EventID == 0 {
		return nil, errors.New("event ID tidak boleh kosong")
	}
	if err := s.transactionRepo.Create(transaction); err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *TransactionService) GetTransactionsByUser(userID string) ([]models.Transaction, error) {
	return s.transactionRepo.FindByUserID(userID)
}

func (s *TransactionService) GetTransactionByID(transactionID int) (*models.Transaction, error) {
	return s.transactionRepo.FindByID(transactionID)
}

func (s *TransactionService) GetTransactionsByEvent(eventID int) ([]models.Transaction, error) {
	return s.transactionRepo.FindByEventID(eventID)
}

func (s *TransactionService) UpdateTransaction(transaction *models.Transaction) error {
	return s.transactionRepo.Update(transaction)
}

func (s *TransactionService) DeleteTransaction(transactionID int) error {
	return s.transactionRepo.Delete(transactionID)
}
