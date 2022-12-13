package repositories

import (
	"github.com/octane77/rova/transactionService/entities"
	"gorm.io/gorm"
)

func NewTransactionRepository(db *gorm.DB) Repository[entities.Transaction] {
	return NewRepository[entities.Transaction](db)
}

//{
//	fir
//	las
//	tx: [
//	{
//		id
//		amount
//		accountId
//		tx_type
//
//	}
//
//	]
//}
