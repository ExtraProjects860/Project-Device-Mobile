package response

import (
	"github.com/gin-gonic/gin"
)

func SendJSON(ctx *gin.Context, code int, payload any) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, payload)
}

func SendSuccess(ctx *gin.Context, code int, payload any) {
	SendJSON(ctx, code, payload)
}

func SendStatus(ctx *gin.Context, code int, message string) {
	SendJSON(ctx, code, Status{
		Message: message,
	})
}

func SendErr(ctx *gin.Context, code int, err error) {
	SendJSON(ctx, code, ErrResponse{
		Error: err.Error(),
	})
}

func SendErrAbort(ctx *gin.Context, code int, err error) {
	ctx.Header("Content-type", "application/json")
	ctx.AbortWithStatusJSON(code, ErrResponse{
		Error: err.Error(),
	})
}

type ErrResponse struct {
	Error string `json:"error"`
}

type Status struct {
	Message string `json:"message"`
}

type TokenResponse struct {
	Access  string `json:"access_token,omitempty"`
	Refresh string `json:"refresh_token,omitempty"`
}
