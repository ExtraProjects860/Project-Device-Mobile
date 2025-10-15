package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"github.com/jackc/pgx/v5/pgconn"
)

func verifyEnterpriseDuplicated(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		if strings.Contains(pgErr.ConstraintName, "name") {
			return errors.New("this name enterprise is registered")
		}
		return err
	}
	return err
}

func (r *PostgresEnterpriseRepository) GetEnterprise(ctx context.Context, id uint) (schemas.Enterprise, error) {
	query := r.db.WithContext(ctx).Model(&schemas.Enterprise{})

	enterprise, err := getByID[schemas.Enterprise](query, id)
	if err != nil {
		return schemas.Enterprise{}, err
	}
	return enterprise, nil
}

func (r *PostgresEnterpriseRepository) CreateEnterprise(ctx context.Context, enterprise *schemas.Enterprise) error {
	err := create(ctx, r.db, enterprise)
	if err != nil {
		return verifyEnterpriseDuplicated(err)
	}

	e, err := r.GetEnterprise(ctx, enterprise.ID)
	if err != nil {
		return err
	}
	*enterprise = e

	return nil
}

func (r *PostgresEnterpriseRepository) UpdateEnterprise(ctx context.Context, id uint, enterprise *schemas.Enterprise) error {
	if err := updateByID(ctx, r.db, enterprise, id); err != nil {
		return verifyEnterpriseDuplicated(err)
	}

	e, err := r.GetEnterprise(ctx, enterprise.ID)
	if err != nil {
		return err
	}
	*enterprise = e

	return nil
}

func (r *PostgresEnterpriseRepository) GetEnterprises(ctx context.Context, itemsPerPage uint, currentPage uint) ([]schemas.Enterprise, uint, uint, error) {
	query := r.db.WithContext(ctx).Model(&schemas.Enterprise{})

	enterprises, totalPages, totalItems, err := getByPagination[schemas.Enterprise](
		query, 
		itemsPerPage, 
		currentPage,
	)
	if err != nil {
		return nil, 0, 0, err
	}

	return enterprises, totalPages, totalItems, err
}
