package usecase

import (
	"context"
	"time"

	"github.com/pramudya3/backend/payment/domain"
)

type userUsecase struct {
	UserRepository           domain.UserRepository
	PaymentAccountRepository domain.PaymentAccountRepository
}

// Create implements domain.UserUsecase.
func (u *userUsecase) Create(ctx context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return u.UserRepository.Insert(ctx, user)
}

// GetById implements domain.UserUsecase.
func (u *userUsecase) GetById(ctx context.Context, id uint64) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	user, err := u.UserRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	payments, err := u.PaymentAccountRepository.FetchPaymentAccount(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	user.PaymentAccounts = payments

	return user, nil
}

func NewUserUsecase(userRepo domain.UserRepository, pRepo domain.PaymentAccountRepository) domain.UserUsecase {
	return &userUsecase{
		UserRepository:           userRepo,
		PaymentAccountRepository: pRepo,
	}
}
