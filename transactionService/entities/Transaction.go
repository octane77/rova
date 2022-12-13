package entities

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID         uint           `gorm:"primary_key; autoIncrement" json:"id"`
	CustomerID uint           `gorm:"not null" json:"customerId"`
	AccountID  uint           `gorm:"not null" json:"accountId"`
	Amount     float64        `gorm:"not null" json:"amount"`
	CreatedAt  time.Time      `gorm:"default: CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt  time.Time      `gorm:"default: CURRENT_TIMESTAMP" json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
