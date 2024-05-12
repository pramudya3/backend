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
func (h *transactionRepository) FetchTransaction(ctx context.Context, userId uint64) ([]*domain.Transaction, error) {
	histories := []*domain.Transaction{}
	res := h.DB.Find(&histories, "user_id = ?", userId)
	if res.Error != nil {
		return nil, res.Error
	}

	return histories, nil
}

func NewTransactionRepository(db *gorm.DB) domain.TransactionRepository {
	return &transactionRepository{
		DB: db,
	}
}
