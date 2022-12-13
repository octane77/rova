package services

import (
	"github.com/mashingan/smapping"
	"github.com/octane77/rova/accountService/dtos"
	"github.com/octane77/rova/accountService/entities"
	"github.com/octane77/rova/accountService/repositories"
)

type AccountService interface {
	CreateCurrentAccount(createAccountDto dtos.CreateAccountDto) (*entities.Account, error)
	GetAccountsByCustomerId(getCustomerAccountsDto dtos.GetCustomerAccountsDto) *[]entities.Account
}

type accountService struct {
	accountRepository repositories.Repository[entities.Account]
}

func (a accountService) CreateCurrentAccount(createAccountDto dtos.CreateAccountDto) (*entities.Account, error) {
	var account entities.Account
	err := smapping.FillStruct(&account, smapping.MapFields(&createAccountDto))
	if err != nil {
		return nil, err
	}
	account.Type = entities.CURRENT_ACCOUNT
	account.Balance = createAccountDto.InitialCredit
	return a.accountRepository.Create(&account)
}

func (a accountService) GetAccountsByCustomerId(getCustomerAccountsDto dtos.GetCustomerAccountsDto) *[]entities.Account {
	return a.accountRepository.FindAllWhere(repositories.FindAllWhereOptions{
		Where: entities.Account{
			CustomerID: getCustomerAccountsDto.CustomerId,
		},
	})
}

func NewAccountService(accountRepository repositories.Repository[entities.Account]) AccountService {
	return accountService{
		accountRepository: accountRepository,
	}
}
