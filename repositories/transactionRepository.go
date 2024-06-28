package repositories

import (
	"gorm.io/gorm"
	"mnc-finance-queue/entity"
)

type TransactionRepository interface {
	Create(tx *gorm.DB, transaction *entity.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) Create(tx *gorm.DB, transaction *entity.Transaction) error {
	return tx.Create(transaction).Error
}
