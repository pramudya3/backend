package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/pramudya3/backend/payment/domain"
	"gorm.io/gorm"
)

type transaction struct {
	paymentAccount  domain.PaymentAccountUsecase
	transactionRepo domain.TransactionRepository
}

// Send implements domain.TransactionUsecase.
func (t *transaction) Send(ctx context.Context, tx *domain.Transaction) (*domain.TransactionResponse, error) {
	paymentAccountSender, err := t.paymentAccount.GetPaymentAccountByType(ctx, tx.UserID, tx.PaymentMethod)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Status = "Failed"
			tx.Description = "Account payment not found"
			return &domain.TransactionResponse{
				Description: tx.Description,
				Status:      tx.Status,
			}, nil
		}
		return nil, err
	}
	if tx.Amount > paymentAccountSender.Balance {
		tx.Status = "Failed"
		tx.Description = "Insufficient balance"
		t.createRecordSender(ctx, tx)
		return &domain.TransactionResponse{
			Description: tx.Description,
			Status:      tx.Status,
		}, nil
	}

	paymentAccountReciever, err := t.paymentAccount.GetPaymentAccountByType(ctx, *tx.ToUserID, tx.PaymentMethod)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Status = "Failed"
			tx.Description = "Account payment not found"
			return &domain.TransactionResponse{
				Description: tx.Description,
				Status:      tx.Status,
			}, nil
		}
		return nil, err
	}

	if paymentAccountSender.ID == paymentAccountReciever.ID {
		tx.Status = "Failed"
		tx.Description = "The sender and receiver cannot be the same"
		return &domain.TransactionResponse{
			Description: tx.Description,
			Status:      tx.Status,
		}, nil
	}

	paymentAccountSender.Balance -= tx.Amount
	paymentAccountReciever.Balance += tx.Amount

	if err := t.paymentAccount.UpdatePaymentAccount(ctx, paymentAccountSender); err != nil {
		return nil, err
	}
	if err := t.paymentAccount.UpdatePaymentAccount(ctx, paymentAccountReciever); err != nil {
		return nil, err
	}

	tx.Status = "Success"
	// insert into payment history table

	tx.Description = fmt.Sprintf("Recieved from user is %d, with ammount %v", tx.UserID, tx.Amount)
	t.createRecordReciever(ctx, tx)

	tx.Description = fmt.Sprintf("Successfully sent to user id %d with ammount %v", *tx.ToUserID, tx.Amount)
	t.createRecordSender(ctx, tx)

	return &domain.TransactionResponse{
		Description: tx.Description,
		Status:      tx.Status,
	}, nil
}

// Withdraw implements domain.TransactionUsecase.
func (t *transaction) Withdraw(ctx context.Context, tx *domain.Transaction) (*domain.TransactionResponse, error) {
	paymentUser, err := t.paymentAccount.GetPaymentAccountByType(ctx, tx.UserID, tx.PaymentMethod)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Status = "Failed"
			tx.Description = "Account payment not found"
			return &domain.TransactionResponse{
				Description: tx.Description,
				Status:      tx.Status,
			}, nil
		}
		return nil, err
	}
	if tx.Amount > paymentUser.Balance {
		tx.Status = "Failed"
		tx.Description = "Insufficient balance"
		t.createRecordSender(ctx, tx)
		return &domain.TransactionResponse{
			Description: tx.Description,
			Status:      tx.Status,
		}, nil
	}

	paymentUser.Balance -= tx.Amount

	if err := t.paymentAccount.UpdatePaymentAccount(ctx, paymentUser); err != nil {
		return nil, err
	}

	tx.Status = "Success"
	tx.Description = fmt.Sprintf("Successful withdrawal with ammount %v", tx.Amount)
	t.createRecordWithdraw(ctx, tx)

	return &domain.TransactionResponse{
		Description: tx.Description,
		Status:      tx.Status,
	}, nil
}

func NewTransactionUsecase(pa domain.PaymentAccountUsecase, txRepo domain.TransactionRepository) domain.TransactionUsecase {
	return &transaction{
		paymentAccount:  pa,
		transactionRepo: txRepo,
	}
}

func (t *transaction) CreateRecord(ctx context.Context, history *domain.Transaction) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return t.transactionRepo.CreateRecord(ctx, history)
}

func (t *transaction) createRecordSender(ctx context.Context, tx *domain.Transaction) {
	if err := t.transactionRepo.CreateRecord(ctx, &domain.Transaction{
		UserID:        tx.UserID,
		FromUserID:    nil,
		ToUserID:      tx.ToUserID,
		Amount:        tx.Amount,
		PaymentMethod: tx.PaymentMethod,
		Description:   tx.Description,
		Status:        tx.Status,
	}); err != nil {
		log.Println(err)
	}
}

func (t *transaction) createRecordReciever(ctx context.Context, tx *domain.Transaction) {
	if err := t.transactionRepo.CreateRecord(ctx, &domain.Transaction{
		UserID:        *tx.ToUserID,
		FromUserID:    &tx.UserID,
		ToUserID:      nil,
		Amount:        tx.Amount,
		PaymentMethod: tx.PaymentMethod,
		Description:   tx.Description,
		Status:        tx.Status,
	}); err != nil {
		log.Println(err)
	}
}

func (t *transaction) createRecordWithdraw(ctx context.Context, tx *domain.Transaction) {
	if err := t.transactionRepo.CreateRecord(ctx, &domain.Transaction{
		UserID:        tx.UserID,
		FromUserID:    nil,
		ToUserID:      nil,
		Amount:        tx.Amount,
		PaymentMethod: tx.PaymentMethod,
		Description:   tx.Description,
		Status:        tx.Status,
	}); err != nil {
		log.Println(err)
	}
}
