package handler

import "github.com/gin-gonic/gin"

// @BasePath /api/v1

// @Summary      Create Promotion
// @Description  Creates a new product promotion
// @Tags         promotions
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /promotions [post]
func CreatePromotion(ctx *gin.Context) {
	sendSuccess(ctx, "Add Promotion Product!")
}

// @Summary      Update Promotion
// @Description  Updates an existing product promotion
// @Tags         promotions
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /promotions/{id} [patch]
func UpdatePromotion(ctx *gin.Context) {
	sendSuccess(ctx, "Update Promotion Product!")
}

// @Summary      Get Promotions
// @Description  Returns all product promotions
// @Tags         promotions
// @Produce      json
// @Success      200 {array} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /promotions [get]
func GetPromotions(ctx *gin.Context) {
	sendSuccess(ctx, "Get Promotions Products!")
}
