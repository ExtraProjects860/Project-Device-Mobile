package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/auth"
	"github.com/gin-gonic/gin"
)

func JWTMiddleware(appCtx *appcontext.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		fmt.Println(tokenStr)
		if strings.HasPrefix(strings.ToLower(tokenStr), "bearer ") {
			tokenStr = strings.TrimSpace(tokenStr[7:])
		}

		claims, err := auth.ValidateAccessToken(tokenStr, appCtx.Env.API.JwtKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("user_id", claims.Sub)
		ctx.Next()
	}
}
