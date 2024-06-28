package repositories

import (
	"gorm.io/gorm"
	"mnc-finance-queue/entity"
)

type UserRepository interface {
	Update(tx *gorm.DB, user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Update(tx *gorm.DB, user *entity.User) error {
	return tx.Save(user).Error
}
