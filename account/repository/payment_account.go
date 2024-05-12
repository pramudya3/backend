package repository

import (
	"context"
	"log"

	"github.com/pramudya3/backend/payment/domain"

	"gorm.io/gorm"
)

type paymentAccountRepository struct {
	DB *gorm.DB
}

// AddNewPayment implements domain.PaymentAccountRepository.
func (p *paymentAccountRepository) AddNewPaymentAccount(ctx context.Context, paymentAccount *domain.PaymentAccount) error {
	if res := p.DB.Create(paymentAccount); res.Error != nil {
		log.Println(res.Error)
		return res.Error
	}
	return nil
}

// DeletePaymentAccount implements domain.PaymentAccountRepository.
func (p *paymentAccountRepository) DeletePaymentAccount(ctx context.Context, paymentId uint64) error {
	if res := p.DB.Delete(&domain.PaymentAccount{}, "id = ?", paymentId); res.Error != nil {
		return res.Error
	}

	return nil
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

// UpdatePaymentAccount implements domain.PaymentAccountRepository.
func (p *paymentAccountRepository) UpdatePaymentAccount(ctx context.Context, payment *domain.PaymentAccount) error {
	res := p.DB.Model(&domain.PaymentAccount{}).Update("balance", payment.Balance).Where("id = ?", payment.ID)
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
