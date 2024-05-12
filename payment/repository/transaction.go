package repository

import (
	"context"

	"github.com/pramudya3/backend/payment/domain"

	"gorm.io/gorm"
)

type transactionRepository struct {
	DB *gorm.DB
}

// FetchHistory implements domain.HistoryRepository.
func (t *transactionRepository) CreateRecord(ctx context.Context, history *domain.Transaction) error {
	if err := t.DB.Create(history).Error; err != nil {
		return err
	}

	return nil
}

func NewTransactionRepository(db *gorm.DB) domain.TransactionRepository {
	return &transactionRepository{
		DB: db,
	}
}
