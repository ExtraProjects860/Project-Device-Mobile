package router

import (
	"github.com/gin-gonic/gin"
)

const basePath string = "/api/v1"

func initMainRoutes(r *gin.Engine) {
	r.Group(basePath)
}
