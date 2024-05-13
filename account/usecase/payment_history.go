package usecase

import (
	"context"
	"time"

	"github.com/pramudya3/backend/payment/domain"
)

type transactionUsecase struct {
	PaymentHistoryRepository domain.TransactionRepository
}

// FetchTransaction implements domain.TransactionUsecase.
func (p *transactionUsecase) FetchTransaction(ctx context.Context, userId uint64) ([]*domain.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return p.PaymentHistoryRepository.FetchTransaction(ctx, userId)
}

func NewTransactionUsecase(phRepo domain.TransactionRepository) domain.TransactionUsecase {
	return &transactionUsecase{
		PaymentHistoryRepository: phRepo,
	}
}
