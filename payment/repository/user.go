package repository

import (
	"context"
	"log"

	"github.com/pramudya3/backend/payment/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

// Fetch implements domain.UserRepository.
func (u *userRepository) Fetch(ctx context.Context) ([]*domain.User, error) {
	users := []*domain.User{}

	if err := u.DB.Table("users").Select("*").Joins("left join payment_accounts on payment_accounts.user_id = users.id").Scan(&users).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

// FindById implements domain.UserRepository.
func (u *userRepository) FindById(ctx context.Context, id uint64) (*domain.User, error) {
	user := domain.User{}
	// if res := u.DB.Table("users").Select("*").Joins("join payment_accounts on payment_accounts.user_id = users.id").Scan(&user); res.Error != nil {
	if res := u.DB.First(&user, "id = ?", id); res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

// Insert implements domain.UserRepository.
func (u *userRepository) Insert(ctx context.Context, user *domain.User) error {
	if res := u.DB.Create(&user); res.Error != nil {
		log.Println(res.Error)
		return res.Error
	}
	return nil
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		DB: db,
	}
}
