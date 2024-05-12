package usecase

import (
	"context"
	"time"

	"github.com/pramudya3/backend/payment/domain"
)

type paymentHistoryUsecase struct {
	PaymentHistoryRepository domain.TransactionRepository
}

// FetchTransaction implements domain.PaymentHistoryUsecase.
func (p *paymentHistoryUsecase) FetchTransaction(ctx context.Context, userId uint64) ([]*domain.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return p.PaymentHistoryRepository.FetchTransaction(ctx, userId)
}

func NewPaymentHistoryUsecase(phRepo domain.TransactionRepository) domain.TransactionUsecase {
	return &paymentHistoryUsecase{
		PaymentHistoryRepository: phRepo,
	}
}
