package usecase

import (
	"context"
	"time"

	"github.com/pramudya3/backend/payment/domain"
)

type paymentAccountUsecase struct {
	PaymentTypeRepository domain.PaymentAccountRepository
}

func (p *paymentAccountUsecase) FetchPaymentAccount(ctx context.Context, accountId uint64) ([]*domain.PaymentAccount, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return p.PaymentTypeRepository.FetchPaymentAccount(ctx, accountId)
}

func (p *paymentAccountUsecase) GetPaymentAccountByType(ctx context.Context, accountId uint64, paymentType string) (*domain.PaymentAccount, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return p.PaymentTypeRepository.GetPaymentAccountByType(ctx, accountId, paymentType)
}

func (p *paymentAccountUsecase) UpdatePaymentAccount(ctx context.Context, payment *domain.PaymentAccount) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return p.PaymentTypeRepository.UpdatePaymentAccount(ctx, payment)
}

func NewPaymentTypeUsecase(par domain.PaymentAccountRepository) domain.PaymentAccountUsecase {
	return &paymentAccountUsecase{
		PaymentTypeRepository: par,
	}
}
