package models

import "time"

type (
	Transaction struct {
		UserID        string
		Type          string
		Amount        float64
		Remarks       string
		BalanceBefore float64
		BalanceAfter  float64
	}

	TopUp struct {
		Amount float64 `json:"amount"`
	}

	Transfer struct {
		RecipientID string  `json:"target_user"`
		Amount      float64 `json:"amount"`
		Remarks     string  `json:"remarks"`
	}

	Payment struct {
		PaymentID     string    `json:"payment_id"`
		Amount        float64   `json:"amount"`
		Remarks       string    `json:"remarks"`
		BalanceBefore float64   `json:"balance_before"`
		BalanceAfter  float64   `json:"balance_after"`
		CreatedDate   time.Time `json:"created_date"`
	}
)
