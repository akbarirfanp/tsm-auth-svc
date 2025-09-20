package middleware

import (
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func AuthMiddleware() http.Middleware {
	return func(ctx http.Context) {
		authHeader := ctx.Request().Header("Authorization", "")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.Response().Json(http.StatusUnauthorized, http.Json{
				"error": "Unauthorized",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		
		secret := []byte(facades.Config().GetString("jwt.secret"))
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
		if err != nil {
			ctx.Response().Json(http.StatusUnauthorized, http.Json{
				"error": "Invalid token",
			})
			return
		}

		ctx.Request().Next()
	}
}
