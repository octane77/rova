package dtos

type GetCustomerAccountsDto struct {
	CustomerId uint `json:"customerId" binding:"required"`
}
