package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/octane77/rova/identityService/common/errors"
	"github.com/octane77/rova/identityService/services"
	"strings"
)

func RequiresAuth(jwtService services.JwtService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken := ctx.GetHeader("Authorization")
		if authToken == "" {
			errors.ThrowUnAuthorizedException(ctx, "user not authorized")
			return
		}
		if strings.Contains(authToken, "Bearer ") {
			authToken = strings.Split(authToken, "Bearer ")[1]
		}
		token, err := jwtService.ValidateToken(authToken)
		if err != nil {
			errors.ThrowUnAuthorizedException(ctx, "token is not valid")
			return
		}
		if token.Valid {
			ctx.Set("userId", token.Claims.(jwt.MapClaims)["userId"])
			ctx.Set("token", authToken)
		} else {
			errors.ThrowUnAuthorizedException(ctx, "token is not valid")
			return
		}
	}
}
