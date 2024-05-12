package domain

import (
	"context"
	"time"
)

type Transaction struct {
	ID            uint64    `json:"id" gorm:"primaryKey"`
	UserID        uint64    `json:"user_id"`
	FromUserID    *uint64   `json:"from_user_id"`
	ToUserID      *uint64   `json:"to_user_id"`
	Amount        uint64    `json:"amount"`
	PaymentMethod string    `json:"payment_type"`
	Description   string    `json:"description"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"craeted_at"`
}

type TransactionUsecase interface {
	FetchTransaction(ctx context.Context, paymentId uint64) ([]*Transaction, error)
}

type TransactionRepository interface {
	FetchTransaction(ctx context.Context, paymentId uint64) ([]*Transaction, error)
}
