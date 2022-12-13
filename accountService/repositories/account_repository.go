package repositories

import (
	"github.com/octane77/rova/accountService/entities"
	"gorm.io/gorm"
)

func NewAccountRepository(db *gorm.DB) Repository[entities.Account] {
	return NewRepository[entities.Account](db)
}
