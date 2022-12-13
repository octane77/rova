package dtos

type CreateAccountDto struct {
	CustomerID    uint    `json:"customerId" binding:"required"`
	InitialCredit float64 `json:"initialCredit"`
}
