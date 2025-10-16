package request

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserAndProductIdQuery(ctx *gin.Context) (uint, uint, error) {
	userID := ctx.Query("user_id")
	if userID == "" {
		return 0, 0, ErrParamIsRequired("user_id", "queryParameter")
	}

	productID := ctx.Query("product_id")
	if productID == "" {
		return 0, 0, ErrParamIsRequired("product_id", "queryParameter")
	}

	parsedUserId, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid user_id: %v", err)
	}

	parsedProductId, err := strconv.ParseUint(productID, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid product_id: %v", err)
	}

	return uint(parsedUserId), uint(parsedProductId), nil
}
