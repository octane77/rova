package repositories

import (
	"github.com/octane77/rova/identityService/entities"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) Repository[entities.User] {
	return NewRepository[entities.User](db)
}
