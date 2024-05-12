package domain

import (
	"context"
	"time"
)

type CreatePaymentAccount struct {
	Type    string `json:"type"`
	Balance uint64 `json:"balance"`
}

type PaymentAccount struct {
	ID        uint64     `json:"id" gorm:"primaryKey"`
	User_ID   uint64     `json:"user_id"`
	Type      string     `json:"type"`
	Balance   uint64     `json:"balance"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type PaymentAccountUsecase interface {
	FetchPaymentAccount(ctx context.Context, accountId uint64) ([]*PaymentAccount, error)
	AddNewPaymentAccount(ctx context.Context, paymentAccount *PaymentAccount) error
	UpdatePaymentAccount(ctx context.Context, payment *PaymentAccount) error
	DeletePaymentAccount(ctx context.Context, paymentId uint64) error
}

type PaymentAccountRepository interface {
	FetchPaymentAccount(ctx context.Context, accountId uint64) ([]*PaymentAccount, error)
	AddNewPaymentAccount(ctx context.Context, paymentAccount *PaymentAccount) error
	UpdatePaymentAccount(ctx context.Context, payment *PaymentAccount) error
	DeletePaymentAccount(ctx context.Context, paymentId uint64) error
}
