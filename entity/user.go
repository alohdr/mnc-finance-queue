package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	FirstName    string
	LastName     string
	PhoneNumber  string `gorm:"unique"`
	Address      string
	PIN          string
	Balance      float64
	RefreshToken string
	CreatedAt    time.Time      `gorm:"column:created_at;type:timestamp" json:"created_date"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:timestamp" json:"updated_date"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}
