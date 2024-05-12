package usecase

import (
	"context"
	"time"

	"github.com/pramudya3/backend/payment/domain"
)

type paymentAccountUsecase struct {
	PaymentTypeRepository domain.PaymentAccountRepository
}

// AddNewPayment implements domain.PaymentTypeUsecase.
func (p *paymentAccountUsecase) AddNewPaymentAccount(ctx context.Context, paymentType *domain.PaymentAccount) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return p.PaymentTypeRepository.AddNewPaymentAccount(ctx, paymentType)
}

// DeletePaymentAccount implements domain.PaymentTypeUsecase.
func (p *paymentAccountUsecase) DeletePaymentAccount(ctx context.Context, paymentId uint64) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return p.PaymentTypeRepository.DeletePaymentAccount(ctx, paymentId)
}

// FetchPaymentAccount implements domain.PaymentTypeUsecase.
func (p *paymentAccountUsecase) FetchPaymentAccount(ctx context.Context, accountId uint64) ([]*domain.PaymentAccount, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return p.PaymentTypeRepository.FetchPaymentAccount(ctx, accountId)
}

// UpdatePaymentAccount implements domain.PaymentTypeUsecase.
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
