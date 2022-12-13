package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/octane77/rova/identityService/controllers"
	"github.com/octane77/rova/identityService/middlewares"
	"github.com/octane77/rova/identityService/services"
)

func UserRouter(r *gin.RouterGroup, userController controllers.UserController, jwtService services.JwtService) *gin.RouterGroup {
	r.POST("/create", userController.CreateUser)
	r.POST("/login", userController.Login)
	r.GET("/profile", middlewares.RequiresAuth(jwtService), userController.GetUserDetails)

	return r
}
