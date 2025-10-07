package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutesApiV1(router *gin.Engine) {
	apiV1 := router.Group("/api/v1")

	registerUserRoutes(apiV1)
	registerProductRoutes(apiV1)
	registerAuthRoutes(apiV1)
	registerWishListRoutes(apiV1)
}
