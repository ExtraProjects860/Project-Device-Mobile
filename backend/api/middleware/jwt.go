package middleware

import (
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/auth"
	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        tokenStr := ctx.GetHeader("Authorization")
        if tokenStr == "" {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
            return
        }

        token, err := auth.ValidateAccessToken(tokenStr)
        if err != nil || !token.Valid {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }

        ctx.Next()
    }
}
