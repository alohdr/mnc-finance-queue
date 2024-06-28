package controllers

import (
	"log"
	"mnc-finance-queue/queue"
)

type TransactionController struct {
	transactionService queue.RabbitDefinition
}

func NewTransactionController(transactionService queue.RabbitDefinition) *TransactionController {
	return &TransactionController{transactionService}
}

func (ctrl *TransactionController) Transfer(queueName, routeKey string) error {
	messages, err := ctrl.transactionService.ConsumeMessage(queueName, routeKey)
	if err != nil {
		log.Fatalf("Failed to start consuming messages: %v", err)
		return err
	}

	for message := range messages {
		log.Printf("\n\nConsumed message from %s : %s\n", queueName, message.Body)
	}

	return nil
}
