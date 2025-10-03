package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO ajustar esse krl, pois as funções estão quase a mesma merda, ou seja, fazer uma que outras vão usar

func sendErr(ctx *gin.Context, code int, payload gin.H) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, payload)
}

func sendSuccess(ctx *gin.Context, payload any) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, payload)
}

func sendStatus(ctx *gin.Context, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{"status": msg})
}

type ErrResponse struct {
	Error string `json:"error"`
}
