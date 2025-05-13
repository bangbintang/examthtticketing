package repository

import (
	"ticketing-konser/internal/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *TransactionRepository) FindByID(id int) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := r.db.First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) FindByUserID(userID string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.db.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) FindByEventID(eventID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.db.Where("event_id = ?", eventID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) Update(transaction *models.Transaction) error {
	return r.db.Save(transaction).Error
}

func (r *TransactionRepository) Delete(id int) error {
	return r.db.Delete(&models.Transaction{}, id).Error
}
