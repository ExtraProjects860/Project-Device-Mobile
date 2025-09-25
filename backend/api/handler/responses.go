package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendErr(ctx *gin.Context, code int, payload gin.H) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, payload)
}

func sendSuccess(ctx *gin.Context, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func sendStatus(ctx *gin.Context, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{"status": msg})
}

type ErrResponse struct {
	Error string `json:"error"`
}
