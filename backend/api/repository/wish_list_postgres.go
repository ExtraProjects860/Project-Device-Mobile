package repository

import (
	"context"
	"time"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

type WishListDTO struct {
	ID        uint         `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Products  []ProductDTO `json:"products"`
}

func makeWishListDTO() *WishListDTO {
	return &WishListDTO{}
}

func (r *postgresWishListRepository) AddInWishList(ctx context.Context, wishlist schemas.WishList) {
	return
}

func (r *postgresWishListRepository) GetItemsWishList(ctx context.Context, itemsPerPage uint, currentPage uint) {
	return
}

func (r *postgresWishListRepository) UpdateWishList(ctx context.Context, id uint, wishlist schemas.WishList) {
	return
}
