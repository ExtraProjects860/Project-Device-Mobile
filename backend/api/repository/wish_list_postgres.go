package repository

import "time"

type WishListPostgres interface {
	AddInWishList()
	UpdateWishList()
	GetItensWishList()
}

type WishListDTO struct {
	ID       uint              `json:"id"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	Products []ProductDTO      `json:"products"`
}
