package dto

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

type WishListMinimalDTO struct {
	UserID    uint `json:"user_id" example:"1"`
	ProductID uint `json:"product_id" example:"1"`
}

func MakeWishListMinimalDTO(entry schemas.WishList) *WishListMinimalDTO {
	return &WishListMinimalDTO{
		UserID:    entry.UserID,
		ProductID: entry.ProductID,
	}
}
