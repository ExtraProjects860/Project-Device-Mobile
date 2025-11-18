package request

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserAndProductIdQuery(ctx *gin.Context) (uint, uint, error) {
	parsedUserId, err := GetIdByToken(ctx)
	if err != nil {
		return 0, 0, err
	}

	productID := ctx.Query("product_id")
	if productID == "" {
		return 0, 0, ErrParamIsRequired("product_id", "queryParameter")
	}

	parsedProductId, err := strconv.ParseUint(productID, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid product_id: %v", err)
	}

	return uint(parsedUserId), uint(parsedProductId), nil
}
