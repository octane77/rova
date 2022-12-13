package dtos

type GetAccountTransactionsDto struct {
	AccountID  uint `json:"accountId"`
	CustomerID uint `json:"customerId"`
}
