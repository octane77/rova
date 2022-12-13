package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/octane77/rova/apiGateway/controllers"
)

func GateWayRouter(r *gin.Engine, c controllers.ApiServiceController) *gin.Engine {
	r.Any("/:service/:version/:route", c.CallService)

	return r
}
