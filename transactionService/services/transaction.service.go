package services

import (
	"github.com/mashingan/smapping"
	"github.com/octane77/rova/transactionService/dtos"
	"github.com/octane77/rova/transactionService/entities"
	"github.com/octane77/rova/transactionService/repositories"
)

type TransactionService interface {
	CreateTransaction(createTransactionDto dtos.CreateTransactionDTO) (*entities.Transaction, error)
	GetTransactionsByAccount(dto dtos.GetCustomerTransactionsDto) *[]entities.Transaction
}

type transactionService struct {
	transactionRepository repositories.Repository[entities.Transaction]
}

func (t transactionService) CreateTransaction(createTransactionDto dtos.CreateTransactionDTO) (*entities.Transaction, error) {
	var tranx entities.Transaction
	err := smapping.FillStruct(&tranx, smapping.MapFields(&createTransactionDto))
	if err != nil {
		return nil, err
	}
	return t.transactionRepository.Create(&tranx)
}
func (t transactionService) GetTransactionsByAccount(dto dtos.GetCustomerTransactionsDto) *[]entities.Transaction {
	return t.transactionRepository.FindAllWhere(repositories.FindAllWhereOptions{
		Where: entities.Transaction{
			CustomerID: dto.CustomerId,
			AccountID:  dto.AccountId,
		},
	})
}

func NewTransactionService(transactionRepository repositories.Repository[entities.Transaction]) TransactionService {
	return transactionService{
		transactionRepository: transactionRepository,
	}
}
