package dtos

type CreateTransactionDTO struct {
	AccountID  uint    `json:"accountId" binding:"required"`
	CustomerID uint    `json:"customerId" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
}
