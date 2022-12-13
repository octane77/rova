package dtos

type GetCustomerTransactionsDto struct {
	CustomerId uint `json:"customerId" binding:"required"`
	AccountId  uint `json:"accountId" binding:"required"`
}
