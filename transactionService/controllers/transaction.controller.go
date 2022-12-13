package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/octane77/rova/transactionService/common"
	"github.com/octane77/rova/transactionService/common/errors"
	"github.com/octane77/rova/transactionService/dtos"
	"github.com/octane77/rova/transactionService/services"
	"github.com/octane77/rova/transactionService/utils"
	"net/http"
)

type TransactionController interface {
	CreateTransaction(ctx *gin.Context)
	GetCustomerTransactions(ctx *gin.Context)
}

type transactionController struct {
	transactionService  services.TransactionService
	gateWayProxyService services.GatewayProxyService
}

func (t transactionController) CreateTransaction(ctx *gin.Context) {
	var createTransactionDto dtos.CreateTransactionDTO
	dtoErr := ctx.ShouldBind(&createTransactionDto)
	if dtoErr != nil {
		errors.ThrowUnprocessableEntityException(ctx, dtoErr.Error())
		return
	}
	createdTransaction, err := t.transactionService.CreateTransaction(createTransactionDto)
	if err != nil {
		errors.ThrowBadRequestException(ctx, err.Error())
		return
	}
	res := common.BuildResponse(true, utils.TRANSACTION_CREATED_MESSAGE, createdTransaction)
	ctx.JSON(http.StatusCreated, res)
}
func (t transactionController) GetCustomerTransactions(ctx *gin.Context) {
	var getCustomerTransactionsDto dtos.GetCustomerTransactionsDto
	dtoErr := ctx.ShouldBind(&getCustomerTransactionsDto)
	if dtoErr != nil {
		errors.ThrowUnprocessableEntityException(ctx, dtoErr.Error())
		return
	}
	transactions := t.transactionService.GetTransactionsByAccount(getCustomerTransactionsDto)
	res := common.BuildResponse(true, utils.FETCHED_TRANSACTIONS, transactions)
	ctx.JSON(http.StatusOK, res)
}

func NewAccountController(transactionService services.TransactionService, gateWayProxyService services.GatewayProxyService) TransactionController {
	return transactionController{
		transactionService:  transactionService,
		gateWayProxyService: gateWayProxyService,
	}
}
