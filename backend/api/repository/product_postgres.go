package repository

import "time"

type ProductPostgres interface {
	CreateProduct()
	GetProducts()
	UpdateProducts()
}

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
