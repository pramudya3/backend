package repository

import (
	"context"

	"github.com/pramudya3/backend/payment/domain"

	"gorm.io/gorm"
)

type paymentAccountRepository struct {
	DB *gorm.DB
}

// FetchPaymentAccount implements domain.PaymentAccountRepository.
func (p *paymentAccountRepository) FetchPaymentAccount(ctx context.Context, userId uint64) ([]*domain.PaymentAccount, error) {
	payments := []*domain.PaymentAccount{}
	res := p.DB.Find(&payments, "user_id = ?", userId)
	if res.Error != nil {
		return nil, res.Error
	}

	return payments, nil
}

func (p *paymentAccountRepository) GetPaymentAccountByType(ctx context.Context, userId uint64, paymentType string) (*domain.PaymentAccount, error) {
	payment := &domain.PaymentAccount{}
	res := p.DB.Where("user_id = ?", userId).Where("type = ?", paymentType).First(&payment)
	if res.Error != nil {
		return nil, res.Error
	}

	return payment, nil
}

// UpdatePaymentAccount implements domain.PaymentAccountRepository.
func (p *paymentAccountRepository) UpdatePaymentAccount(ctx context.Context, payment *domain.PaymentAccount) error {
	res := p.DB.Model(&domain.PaymentAccount{}).Where("id = ?", payment.ID).Update("balance", payment.Balance)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func NewPaymentAccountRepository(db *gorm.DB) domain.PaymentAccountRepository {
	return &paymentAccountRepository{
		DB: db,
	}
}
