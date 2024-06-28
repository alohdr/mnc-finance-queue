package services

import (
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"mnc-finance-queue/entity"
	"mnc-finance-queue/repositories"
	"mnc-finance-queue/utils/errorMessage"
)

type TransactionService interface {
	Transfer(obj []byte) error
}

type transactionService struct {
	db                    *gorm.DB
	transactionRepository repositories.TransactionRepository
	userRepository        repositories.UserRepository
}

func NewTransactionService(db *gorm.DB, transactionRepo repositories.TransactionRepository, userRepo repositories.UserRepository) TransactionService {
	return &transactionService{db, transactionRepo, userRepo}
}

func (s *transactionService) Transfer(obj []byte) error {

	var model entity.UserTransaction

	err := json.Unmarshal(obj, &model)
	if err != nil {
		log.Println(err)
		return errorMessage.ErrInternalServerError
	}

	tx := s.db.Begin()

	model.RecipientObj.Balance += model.TransactionObj.Amount
	if err := s.userRepository.Update(tx, &model.RecipientObj); err != nil {
		log.Println(err)
		return errorMessage.ErrInternalServerError
	}

	model.UserObj.Balance -= model.TransactionObj.Amount
	if err := s.userRepository.Update(tx, &model.UserObj); err != nil {
		log.Println(err)
		return errorMessage.ErrInternalServerError
	}

	if err := s.transactionRepository.Create(tx, &model.TransactionObj); err != nil {
		log.Println(err)
		return errorMessage.ErrInternalServerError
	}

	return nil
}
