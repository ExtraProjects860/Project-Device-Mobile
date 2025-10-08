package dto

import "github.com/ExtraProjects860/Project-Device-Mobile/schemas"

type WishListDTO struct {
	UserID    uint         `json:"id"`
	ItemCount int          `json:"item_count"`
	Products  []ProductDTO `json:"products"`
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
