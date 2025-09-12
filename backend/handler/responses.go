package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendErr(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": msg,
	})
}

func sendSuccess(ctx *gin.Context, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

type ErrReponse struct {
	Message string `json:"message"`
}
