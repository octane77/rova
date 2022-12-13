package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/octane77/rova/accountService/controllers"
)

func AccountRouter(r *gin.RouterGroup, accountController controllers.AccountController) *gin.RouterGroup {
	r.POST("/create", accountController.CreateCurrentAccount)
	r.POST("/get-all-by-customer-id", accountController.GetCustomerAccounts)

	return r
}
