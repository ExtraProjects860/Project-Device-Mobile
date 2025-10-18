package middleware

import (
	"errors"
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/enum"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/response"
	"github.com/ExtraProjects860/Project-Device-Mobile/service"
	"github.com/gin-gonic/gin"
)

func AdminPermission(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uidRaw, exists := ctx.Get("user_id")
		if !exists {
			logger.Error("user id not found in token")
			response.SendErrAbort(ctx, http.StatusUnauthorized, errors.New("user id not found in token"))
			return
		}

		uid, ok := uidRaw.(uint)
		if !ok {
			logger.Error("invalid user id type")
			response.SendErrAbort(ctx, http.StatusInternalServerError, errors.New("invalid user id type"))
			return
		}

		userService := service.GetUserService(appCtx)
		user, err := userService.Get(ctx, uid)
		if err != nil {
			logger.Error(err)
			response.SendErrAbort(ctx, http.StatusInternalServerError, errors.New("error to process get user"))
			return
		}

		if user.Role != enum.Admin.String() && user.Role != enum.SuperAdmin.String() {
			logger.Error("user don't have permission ADMIN or SUPERADMIN to access")
			response.SendErrAbort(ctx, http.StatusForbidden, errors.New("user don't have permission ADMIN or SUPERADMIN to access"))
			return 
		}

		ctx.Next()
	}
}
