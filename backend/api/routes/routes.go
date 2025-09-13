package routes

import (
	"github.com/gin-gonic/gin"
)

func InitMainRoutes(r *gin.Engine) {
	const basePath string = "/api/v1"

	r.Group(basePath)
}
