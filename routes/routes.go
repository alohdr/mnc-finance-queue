package routes

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"mnc-finance-queue/config"
	"mnc-finance-queue/controllers"
	"mnc-finance-queue/queue"
	"mnc-finance-queue/repositories"
	"mnc-finance-queue/services"
	"mnc-finance-queue/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SetupRoutes() error {
	server := &http.Server{
		Addr: fmt.Sprintf("%v:%v", os.Getenv("HOST"), os.Getenv("PORT")),
	}

	db := config.SetupDatabase()

	mq := config.SetUpRabbitMQ()

	userRepo := repositories.NewUserRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)

	transactionService := services.NewTransactionService(db, transactionRepo, userRepo)

	rabbit := queue.NewRabbitService(mq, transactionService)

	transactionController := controllers.NewTransactionController(rabbit)

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")

	// Declare queues with routing keys
	queues := map[string]string{
		utils.EventMncTransfer: utils.RouteMncTransfer,
	}
	g, _ := errgroup.WithContext(context.Background())

	for queueName, routeKey := range queues {
		queueName := queueName
		routeKey := routeKey
		g.Go(func() error {
			return transactionController.Transfer(queueName, routeKey)
			// Bind the queue to the exchange with the specified routing key
		})
	}

	errGo := g.Wait()
	if errGo != nil {
		log.Println("[Server][Error]: ", errGo)
		log.Fatal(errGo)
	}

	return handleShutdown(server)
}

func handleShutdown(server *http.Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	if err = server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
	return err
}
