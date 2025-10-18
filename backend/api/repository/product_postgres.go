package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"github.com/jackc/pgx/v5/pgconn"
)

func verifyProductDuplicated(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		if strings.Contains(pgErr.ConstraintName, "name") {
			return errors.New("this name product is registered")
		}
		return err
	}
	return err
}

func (r *PostgresProductRepository) GetProduct(ctx context.Context, id uint) (schemas.Product, error) {
	query := r.db.WithContext(ctx).Model(&schemas.Product{})

	product, err := getByID[schemas.Product](query, id)
	if err != nil {
		return schemas.Product{}, err
	}
	return product, nil
}

func (r *PostgresProductRepository) CreateProduct(ctx context.Context, product *schemas.Product) error {
	err := create(ctx, r.db, product)
	if err != nil {
		return verifyProductDuplicated(err)
	}

	u, err := r.GetProduct(ctx, product.ID)
	if err != nil {
		return err
	}
	*product = u

	return nil
}

func (r *PostgresProductRepository) GetProducts(ctx context.Context, itemsPerPage uint, currentPage uint) ([]schemas.Product, uint, uint, error) {
	query := r.db.WithContext(ctx).Model(&schemas.Product{})

	products, totalPages, totalItems, err := getByPagination[schemas.Product](
		query, 
		itemsPerPage, 
		currentPage,
	)
	if err != nil {
		return nil, 0, 0, err
	}

	return products, totalPages, totalItems, err
}

func (r *PostgresProductRepository) UpdateProducts(ctx context.Context, id uint, product *schemas.Product) error {
	if err := updateByID(ctx, r.db, product, id); err != nil {
		return verifyProductDuplicated(err)
	}

	p, err := r.GetProduct(ctx, id)
	if err != nil {
		return err
	}

	*product = p

	return nil
}
