package entities

import (
	"gorm.io/gorm"
	"time"
)

type AccountType string

const (
	CURRENT_ACCOUNT = "current"
)

type Account struct {
	ID         uint           `gorm:"primary_key; autoIncrement" json:"id"`
	Type       string         `gorm:"size: 255; enum('current', 'savings', 'fixed'); not null" json:"type"`
	CustomerID uint           `gorm:"not null" json:"customerId"`
	Balance    float64        `gorm:"not null" json:"balance"`
	CreatedAt  time.Time      `gorm:"default: CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt  time.Time      `gorm:"default: CURRENT_TIMESTAMP" json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
