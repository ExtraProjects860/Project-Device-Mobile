package dto

import (
	"time"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

type ProductDTO struct {
	ID                 uint      `json:"id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	Value              float64   `json:"value"`
	Quantity           int       `json:"quantity"`
	IsPromotionAvaible bool      `json:"is_promotion_avaible"`
	Discount           *float64  `json:"discount,omitempty"`
	PhotoUrl           *string   `json:"photo_url,omitempty"`
	IsAvaible          bool      `json:"is_avaible"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func MakeProductOutput(product schemas.Product) *ProductDTO {
	return &ProductDTO{
		ID:                 product.ID,
		Name:               product.Name,
		Description:        product.Description,
		Value:              product.Value,
		Quantity:           product.Quantity,
		IsPromotionAvaible: product.IsPromotionAvaible,
		Discount:           product.Discount,
		PhotoUrl:           product.PhotoUrl,
		IsAvaible:          product.IsAvaible,
		CreatedAt:          product.CreatedAt,
		UpdatedAt:          product.UpdatedAt,
	}
}
