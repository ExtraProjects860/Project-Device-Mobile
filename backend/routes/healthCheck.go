package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initPingRoute(r *gin.Engine) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong!!!",
		})
	})
}
