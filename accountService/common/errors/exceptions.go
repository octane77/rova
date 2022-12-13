package errors

import (
	"github.com/gin-gonic/gin"
	"github.com/octane77/rova/accountService/common"
	"net/http"
)

func ThrowBadRequestException(ctx *gin.Context, err string) {
	res := common.BuildErrorResponse("Bad Request", err)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	return
}

func ThrowUnprocessableEntityException(ctx *gin.Context, err string) {
	res := common.BuildErrorResponse("Validator Error", err)
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
	return
}

func ThrowUnAuthorizedException(ctx *gin.Context, err string) {
	res := common.BuildErrorResponse("UnAuthorized", err)
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
	return
}

func ThrowForbiddenException(ctx *gin.Context, err string) {
	res := common.BuildErrorResponse("Forbidden", err)
	ctx.AbortWithStatusJSON(http.StatusForbidden, res)
	return
}

func ThrowInternalServerError(ctx *gin.Context, err string) {
	res := common.BuildErrorResponse("Internal Server Error", err)
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
	return
}

func ThrowNotFoundException(ctx *gin.Context) {
	res := common.BuildErrorResponse("NotFound", "route not found")
	ctx.AbortWithStatusJSON(http.StatusNotFound, res)
	return
}
