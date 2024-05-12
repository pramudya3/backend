package domain

import (
	"context"
	"time"
)

type TransactionSend struct {
	ToUserID      uint64 `json:"to_user_id"`
	Amount        uint64 `json:"amount"`
	PaymentMethod string `json:"payment_method"`
}

type TransactionWithdraw struct {
	Amount        uint64 `json:"amount"`
	PaymentMethod string `json:"payment_method"`
}

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

type TransactionResponse struct {
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TransactionUsecase interface {
	Send(ctx context.Context, paymentAccount *Transaction) (*TransactionResponse, error)
	Withdraw(ctx context.Context, paymentAccount *Transaction) (*TransactionResponse, error)
	CreateRecord(ctx context.Context, history *Transaction) error
}

type TransactionRepository interface {
	CreateRecord(ctx context.Context, history *Transaction) error
}
