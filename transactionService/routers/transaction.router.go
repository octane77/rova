package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/octane77/rova/transactionService/controllers"
)

func TransactionRouter(r *gin.RouterGroup, transactionController controllers.TransactionController) *gin.RouterGroup {
	r.POST("/create", transactionController.CreateTransaction)
	r.POST("/get-all-by-account", transactionController.GetCustomerTransactions)

	return r
}
