package routes

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/gin-gonic/gin"
)

func InitRoutesApiV1(router *gin.Engine, appCtx *appcontext.AppContext) {
	apiV1 := router.Group("/api/v1")

	registerUserRoutes(apiV1, appCtx)
	registerProductRoutes(apiV1, appCtx)
	registerAuthRoutes(apiV1, appCtx)
	registerWishListRoutes(apiV1, appCtx)
	registerEnterpriseRoutes(apiV1, appCtx)
	registerRoleRoutes(apiV1, appCtx)
}
