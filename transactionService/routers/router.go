package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/octane77/rova/transactionService/common/errors"
	"github.com/octane77/rova/transactionService/controllers"
	"net/http"
)

func NewRouter(transactionController controllers.TransactionController) *gin.Engine {
	r := gin.Default()
	// HEALTH CHECK
	r.GET("/", func(c *gin.Context) {
		//time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "CURRENT ACCOUNT SERVICE RUNNING")

	})
	v1 := r.Group("/transaction/v1")
	TransactionRouter(v1, transactionController)

	r.NoRoute(func(c *gin.Context) {
		errors.ThrowNotFoundException(c)
	})
	return r
}
