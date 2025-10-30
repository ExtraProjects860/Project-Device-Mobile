package middleware

import (
	"errors"
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/response"
	"github.com/ExtraProjects860/Project-Device-Mobile/utils"
	"github.com/gin-gonic/gin"
)

var (
	ImageRequired = func(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
		return ImageMiddleware(appCtx, logger, true)
	}
	ImageOptional = func(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
		return ImageMiddleware(appCtx, logger, false)
	}
)

func SetAllowedImageSize(router *gin.Engine) {
	router.MaxMultipartMemory = 8 << 20
}

func ImageMiddleware(appCtx *appcontext.AppContext, logger *config.Logger, required bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, err := ctx.FormFile("image")
		// TODO melhorar msg erro aqui depois
		if err != nil {
			if required {
				logger.Error("param image 'image' missing")
				response.SendErrAbort(ctx, http.StatusBadRequest, errors.New("param image 'photo_url' missing"))
				return
			}
			ctx.Next()
			return
		}

		if !utils.IsImage(file.Header.Get("Content-Type")) {
			// TODO verificar como fazer um for para exibir imagens perimitidas
			logger.Error("Invalid file type. Only images are allowed. Tyṕes: ")
			response.SendErrAbort(ctx, http.StatusBadRequest, errors.New("invalid file type. Only images are allowed. Tyṕes: "))
			return
		}

		ctx.Set("image", file)
		ctx.Next()
	}
}
