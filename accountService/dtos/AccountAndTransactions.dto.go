package dtos

type AccountsAndTransactionsDto struct {
	ID           uint        `json:"id"`
	Type         string      `json:"type"`
	CustomerID   uint        `json:"customerId"`
	Balance      float64     `json:"balance"`
	Transactions interface{} `json:"transactions"`
}
