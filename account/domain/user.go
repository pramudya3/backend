package domain

import (
	"context"
	"time"
)

type Login struct {
	ID       uint64 `json:"id"`
	Password string `json:"password"`
}

type Signup struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID              uint64            `json:"id" gorm:"primaryKey"`
	Name            string            `json:"name"`
	Username        string            `json:"username" gorm:"unique"`
	Email           string            `json:"email" gorm:"unique"`
	Password        string            `json:"password"`
	CreatedAt       time.Time         `json:"created_at"`
	Updated         *time.Time        `json:"updated_at"`
	PaymentAccounts []*PaymentAccount `json:"payment_accounts,omitempty"`
	PaymentHistory  []*Transaction    `json:"payment_history,omitempty" gorm:"foreignKey:FromUserID"`
}

type UserUsecase interface {
	Create(ctx context.Context, user *User) error
	GetById(ctx context.Context, id uint64) (*User, error)
}

type UserRepository interface {
	Insert(ctx context.Context, user *User) error
	FindById(ctx context.Context, id uint64) (*User, error)
}
