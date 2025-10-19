package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/auth"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/response"
	"github.com/gin-gonic/gin"
)

func JWTMiddleware(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" {
			logger.Error("missing token")
			response.SendErrAbort(ctx, http.StatusBadRequest, errors.New("missing token"))
			return
		}

		if strings.HasPrefix(strings.ToLower(tokenStr), "bearer ") {
			tokenStr = strings.TrimSpace(tokenStr[7:])
		}

		claims, err := auth.ValidateAccessToken(tokenStr, appCtx.Env.API.JwtKey)
		if err != nil {
			logger.Error(err.Error())
			response.SendErrAbort(ctx, http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		ctx.Set("user_id", claims.Sub)
		ctx.Next()
	}
}
