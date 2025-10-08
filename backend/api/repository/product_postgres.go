package repository

import (
	"context"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

func (r *PostgresProductRepository) CreateProduct(ctx context.Context, product schemas.Product) (schemas.Product, error) {
	err := create(ctx, r.db, &product)
	if err != nil {
		return schemas.Product{}, err
	}
	return product, nil
}

func (r *PostgresProductRepository) GetProducts(ctx context.Context, itemsPerPage uint, currentPage uint) ([]schemas.Product, uint, uint, error) {
	query := r.db.WithContext(ctx).Model(&schemas.Product{})

	products, totalPages, totalItems, err := getByPagination[schemas.Product](query, itemsPerPage, currentPage)
	if err != nil {
		return nil, 0, 0, err
	}

	return products, totalPages, totalItems, err
}

func (r *PostgresProductRepository) UpdateProducts(ctx context.Context, id uint, product schemas.Product) (schemas.Product, error) {
	panic("Not implemented")
}
