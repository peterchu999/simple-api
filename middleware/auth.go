package middleware

import (
	"net/http"
	auth "peterchu999/simple-api/utils/auth"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := auth.ValidateToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
