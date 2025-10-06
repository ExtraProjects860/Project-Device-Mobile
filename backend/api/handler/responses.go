package handler

import (
	"github.com/gin-gonic/gin"
)

func sendJSON(ctx *gin.Context, code int, payload any) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, payload)
}

func sendSuccess(ctx *gin.Context, code int, payload any) {
	sendJSON(ctx, code, payload)
}

func sendStatus(ctx *gin.Context, code int, message string) {
	sendJSON(ctx, code, Status{
		Message: message,
	})
}

func sendErr(ctx *gin.Context, code int, err error) {
	sendJSON(ctx, code, ErrResponse{
		Error: err.Error(),
	})
}

type ErrResponse struct {
	Error string `json:"error"`
}

type Status struct {
	Message string `json:"message"`
}
