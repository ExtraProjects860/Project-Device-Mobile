package dto

import (
	"time"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

type ProductDTO struct {
	ID                 uint      `json:"id" example:"100"`
	Name               string    `json:"name" example:"Smartphone X"`
	Description        string    `json:"description" example:"Smartphone de última geração"`
	Value              float64   `json:"value" example:"2999.99"`
	Quantity           int       `json:"quantity" example:"50"`
	IsPromotionAvaible *bool     `json:"is_promotion_avaible" example:"true"`
	Discount           *float64  `json:"discount,omitempty" example:"10.5"`
	PhotoUrl           *string   `json:"photo_url,omitempty" example:"https://cdn.exemplo.com/produtos/smartphone-x.jpg"`
	IsAvaible          *bool     `json:"is_avaible" example:"true"`
	CreatedAt          time.Time `json:"created_at" example:"2025-10-12T20:00:00Z"`
	UpdatedAt          time.Time `json:"updated_at" example:"2025-10-12T21:00:00Z"`
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
