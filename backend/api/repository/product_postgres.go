package repository

import (
	"context"
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

func makeProductOutput(product schemas.Product) *ProductDTO {
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

func (r *postgresProductRepository) CreateProduct(ctx context.Context, product schemas.Product) {
	return
}

func (r *postgresProductRepository) GetProducts(ctx context.Context, itemsPerPage uint, currentPage uint) (PaginationDTO, error) {
	query := r.db.WithContext(ctx).Model(&schemas.Product{})
<<<<<<< HEAD
	paginationOffset, totalPages := pagination(query, itemsPerPage, currentPage)
=======
	paginationOffset, totalPages, lengthItems := pagination(query, itemsPerPage, currentPage)
>>>>>>> dev

	var productsEntries []schemas.Product
	err := query.
		Limit(int(itemsPerPage)).
		Offset(int(paginationOffset)).
		Find(&productsEntries).Error
	if err != nil {
		logger.Errorf("%v", err)
		return PaginationDTO{}, err
	}

	var productsDTO []ProductDTO
	for _, product := range productsEntries {
		productsDTO = append(productsDTO, *makeProductOutput(product))
	}

<<<<<<< HEAD
	return PaginationDTO{Data: productsDTO, CurrentPage: currentPage, TotalPages: totalPages}, err
=======
	return PaginationDTO{Data: productsDTO, CurrentPage: currentPage, TotalPages: totalPages, TotalItems: lengthItems}, err
>>>>>>> dev
}

func (r *postgresProductRepository) UpdateProducts(ctx context.Context, id uint) {
	return
}
