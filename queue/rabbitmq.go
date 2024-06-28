// queue/rabbitmq.go
package queue

import (
	"log"
	"mnc-finance-queue/config"
	"mnc-finance-queue/services"
	"mnc-finance-queue/utils"
	"os"
)

type RabbitDefinition interface {
	ConsumeMessage(queueName, routeKey string) (<-chan *Message, error)
}
type RabbitService struct {
	rabbit             config.RabbitMQ
	transactionService services.TransactionService
}

func NewRabbitService(rabbit config.RabbitMQ, transactionService services.TransactionService) RabbitDefinition {
	return RabbitService{
		rabbit:             rabbit,
		transactionService: transactionService,
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

func (p RabbitService) ConsumeMessage(queueName, routeKey string) (<-chan *Message, error) {
	// Consume messages from the specified queue
	_, err := p.rabbit.Channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	failOnError(err, "Failed to register a consumer")

	// Bind the queue to the exchange with a wildcard routing key
	err = p.rabbit.Channel.QueueBind(
		queueName,
		routeKey,
		os.Getenv("RabbitExchange"),
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	// Consume messages from the unique queue
	msgs, err := p.rabbit.Channel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	messages := make(chan *Message)

	go func() {
		for msg := range msgs {
			message := &Message{
				Body: string(msg.Body),
			}
			messages <- message

			switch queueName {
			case utils.EventMncTransfer:
				err = p.transactionService.Transfer(msg.Body)
				if err != nil {
					log.Printf("failed create log err brigate: %s", err.Error())
				}

			}

		}
	}()

	return messages, nil
}
