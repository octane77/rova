package services

import (
	"errors"
	"github.com/mashingan/smapping"
	"github.com/octane77/rova/accountService/dtos"
	"github.com/octane77/rova/accountService/entities"
	"github.com/octane77/rova/accountService/utils"
	"os"
)

type GatewayProxyService interface {
	CreateTransaction(account *entities.Account, amount float64) (*interface{}, error)
	GetAccountsTransactions(accounts []entities.Account) (*[]dtos.AccountsAndTransactionsDto, error)
}

type gatewayProxyService struct {
}

func (g gatewayProxyService) CreateTransaction(account *entities.Account, amount float64) (*interface{}, error) {
	url := os.Getenv("TRANSACTION_SERVICE_URL")
	if url == "" {
		return nil, errors.New("transaction service url not found")
	}
	body := dtos.CreateTransactionDTO{
		AccountID:  account.ID,
		CustomerID: account.CustomerID,
		Amount:     amount,
	}
	return utils.MakeApiRequest(url+"/create", body, utils.API_REQ_METHOD_POST)
}

func (g gatewayProxyService) GetAccountsTransactions(accounts []entities.Account) (*[]dtos.AccountsAndTransactionsDto, error) {
	var accountsAndTransactions []dtos.AccountsAndTransactionsDto
	url := os.Getenv("TRANSACTION_SERVICE_URL")
	if url == "" {
		return nil, errors.New("transaction service url not found")
	}
	for _, v := range accounts {
		var accAndTranx dtos.AccountsAndTransactionsDto
		err := smapping.FillStruct(&accAndTranx, smapping.MapFields(&v))
		if err != nil {
			return nil, err
		}

		body := dtos.GetAccountTransactionsDto{
			CustomerID: v.CustomerID,
			AccountID:  v.ID,
		}

		res, err := utils.MakeApiRequest(url+"/get-all-by-account", body, utils.API_REQ_METHOD_POST)
		if err != nil {
			return nil, err
		}
		accAndTranx.Transactions = (*res).(map[string]interface{})["data"]

		accountsAndTransactions = append(accountsAndTransactions, accAndTranx)
	}
	return &accountsAndTransactions, nil
}

func NewGatewayProxyService() GatewayProxyService {
	return gatewayProxyService{}
}
