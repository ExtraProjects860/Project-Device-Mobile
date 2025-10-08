package repository

import (
	"context"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

func (r *PostgresEnterpriseRepository) CreateEnterprise(ctx context.Context, enterprise schemas.Enterprise) (schemas.Enterprise, error) {
	err := create(ctx, r.db, &enterprise)
	if err != nil {
		return schemas.Enterprise{}, err
	}
	return enterprise, nil
}

func (r *PostgresEnterpriseRepository) GetEnterprises(ctx context.Context, itemsPerPage uint, currentPage uint) ([]schemas.Enterprise, uint, uint, error) {
	query := r.db.WithContext(ctx).Model(&schemas.Enterprise{})

	enterprises, totalPages, totalItems, err := getByPagination[schemas.Enterprise](query, itemsPerPage, currentPage)
	if err != nil {
		return nil, 0, 0, err
	}

	return enterprises, totalPages, totalItems, err
}

func (r *PostgresEnterpriseRepository) UpdateEnterprise(ctx context.Context, id uint, enterprise schemas.Enterprise) (schemas.Enterprise, error) {
	panic("Not implemented")
}
