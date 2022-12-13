package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/octane77/rova/apiGateway/common/errors"
	"github.com/octane77/rova/apiGateway/dtos"
	"github.com/octane77/rova/apiGateway/services"
	"github.com/octane77/rova/apiGateway/utils"
	"net/http"
)

type ApiServiceController interface {
	CallService(ctx *gin.Context)
}

type apiServiceController struct {
	accountService  services.AccountService
	identityService services.IdentityService
}

func (a apiServiceController) CallService(ctx *gin.Context) {
	var body interface{}

	method := ctx.Request.Method
	if method != utils.API_REQ_METHOD_GET {
		err := ctx.ShouldBind(&body)
		if err != nil {
			errors.ThrowBadRequestException(ctx, err.Error())
			return
		}
	}
	fmt.Println(method)
	dto := dtos.CallServiceApiDto{
		Route:  ctx.Param("route"),
		Body:   body,
		Method: method,
		Path:   ctx.Request.URL.Path,
		Token:  ctx.GetHeader("Authorization"),
	}
	service := ctx.Param("service")
	switch service {
	case "account":
		res, err := a.accountService.HandleRequest(dto)
		if err != nil {
			errors.ThrowBadRequestException(ctx, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, res)
		break
	case "user":
		res, err := a.identityService.HandleRequest(dto)
		if err != nil {
			errors.ThrowBadRequestException(ctx, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, res)
		break
	default:
		errors.ThrowNotFoundException(ctx)
		return
	}

}

func NewApiServiceController(accountService services.AccountService, identityService services.IdentityService) ApiServiceController {
	return apiServiceController{accountService: accountService, identityService: identityService}
}
