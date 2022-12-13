package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/octane77/rova/apiGateway/common/errors"
	"github.com/octane77/rova/apiGateway/controllers"
	"net/http"
)

func NewRouter(c controllers.ApiServiceController) *gin.Engine {
	r := gin.Default()
	// HEALTH CHECK
	r.GET("/", func(c *gin.Context) {
		//time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "CURRENT ACCOUNT SERVICE RUNNING")

	})

	GateWayRouter(r, c)

	r.NoRoute(func(c *gin.Context) {
		errors.ThrowNotFoundException(c)
	})
	return r
}
