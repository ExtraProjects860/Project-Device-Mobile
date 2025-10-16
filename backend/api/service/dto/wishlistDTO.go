package dto

import (
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

type WishListDTO struct {
	UserID    uint         `json:"id" example:"1"`
	ItemCount int          `json:"item_count" example:"1"`
	Products  []ProductDTO `json:"products"`
}

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

func MakeWishListOutput(wishListEntries []schemas.WishList, userID uint) *WishListDTO {
	productsDTO := make([]ProductDTO, 0, len(wishListEntries))
	for _, entry := range wishListEntries {
		if entry.Product.ID == 0 {
			continue
		}

		product := ProductDTO{
			ID:          entry.Product.ID,
			Name:        entry.Product.Name,
			Description: entry.Product.Description,
			Value:       entry.Product.Value,
			PhotoUrl:    entry.Product.PhotoUrl,
		}
		productsDTO = append(productsDTO, product)
	}

	return &WishListDTO{
		UserID:    userID,
		ItemCount: len(productsDTO),
		Products:  productsDTO,
	}
}
