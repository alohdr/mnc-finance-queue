package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID            uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"transfer_id"`
	UserID        uuid.UUID      `gorm:"column:user_id;type:uuid" json:"-"`
	RecipientID   uuid.UUID      `gorm:"column:recipient_id;type:uuid" json:"-"`
	Type          string         `gorm:"column:type" json:"-"`
	Amount        float64        `gorm:"column:amount" json:"amount"`
	Remarks       string         `gorm:"column:remarks" json:"remarks"`
	BalanceBefore float64        `gorm:"column:balance_before" json:"balance_before"`
	BalanceAfter  float64        `gorm:"column:balance_after" json:"balance_after"`
	Status        string         `gorm:"column:status" json:"status"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:timestamp" json:"created_date"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"-"`
}

type UserTransaction struct {
	UserObj        User
	RecipientObj   User
	TransactionObj Transaction
}
