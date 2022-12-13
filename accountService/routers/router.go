package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/octane77/rova/accountService/common/errors"
	"github.com/octane77/rova/accountService/controllers"
	"net/http"
)

func NewRouter(accountController controllers.AccountController) *gin.Engine {
	r := gin.Default()
	// HEALTH CHECK
	r.GET("/", func(c *gin.Context) {
		//time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "CURRENT ACCOUNT SERVICE RUNNING")

	})
	v1 := r.Group("/account/v1")
	AccountRouter(v1, accountController)

	r.NoRoute(func(c *gin.Context) {
		errors.ThrowNotFoundException(c)
	})
	return r
}
