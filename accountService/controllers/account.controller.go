package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/octane77/rova/accountService/common"
	"github.com/octane77/rova/accountService/common/errors"
	"github.com/octane77/rova/accountService/dtos"
	"github.com/octane77/rova/accountService/services"
	"github.com/octane77/rova/accountService/utils"
	"net/http"
)

type AccountController interface {
	CreateCurrentAccount(ctx *gin.Context)
	GetCustomerAccounts(ctx *gin.Context)
}

type accountController struct {
	accountService      services.AccountService
	gateWayProxyService services.GatewayProxyService
}

func (a accountController) CreateCurrentAccount(ctx *gin.Context) {
	var createAccountDto dtos.CreateAccountDto
	dtoErr := ctx.ShouldBind(&createAccountDto)
	if dtoErr != nil {
		errors.ThrowUnprocessableEntityException(ctx, dtoErr.Error())
		return
	}
	createdAccount, err := a.accountService.CreateCurrentAccount(createAccountDto)
	if err != nil {
		errors.ThrowBadRequestException(ctx, err.Error())
		return
	}
	if createAccountDto.InitialCredit > 0 {
		_, err = a.gateWayProxyService.CreateTransaction(createdAccount, createAccountDto.InitialCredit)
		if err != nil {
			errors.ThrowBadRequestException(ctx, err.Error())
			return
		}
	}

	res := common.BuildResponse(true, utils.ACCOUNT_CREATED_MESSAGE, createdAccount)
	ctx.JSON(http.StatusCreated, res)
}

func (a accountController) GetCustomerAccounts(ctx *gin.Context) {
	var getCustomerAccountsDto dtos.GetCustomerAccountsDto
	dtoErr := ctx.ShouldBind(&getCustomerAccountsDto)
	if dtoErr != nil {
		errors.ThrowUnprocessableEntityException(ctx, dtoErr.Error())
		return
	}
	accounts := a.accountService.GetAccountsByCustomerId(getCustomerAccountsDto)
	accAndTranx, err := a.gateWayProxyService.GetAccountsTransactions(*accounts)
	if err != nil {
		errors.ThrowBadRequestException(ctx, err.Error())
		return
	}
	res := common.BuildResponse(true, utils.ACCOUNTS_FETCHED, accAndTranx)
	ctx.JSON(http.StatusCreated, res)
}

func NewAccountController(accountService services.AccountService, gateWayProxyService services.GatewayProxyService) AccountController {
	return accountController{
		accountService:      accountService,
		gateWayProxyService: gateWayProxyService,
	}
}
