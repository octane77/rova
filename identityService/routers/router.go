package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/octane77/rova/identityService/common/errors"
	"github.com/octane77/rova/identityService/controllers"
	"github.com/octane77/rova/identityService/services"
	"net/http"
)

func NewRouter(userController controllers.UserController, jwtService services.JwtService) *gin.Engine {
	r := gin.Default()
	// HEALTH CHECK
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "IDENTITY SERVICE RUNNING")

	})
	v1 := r.Group("/user/v1")
	UserRouter(v1, userController, jwtService)

	r.NoRoute(func(c *gin.Context) {
		errors.ThrowNotFoundException(c)
	})
	return r

}
