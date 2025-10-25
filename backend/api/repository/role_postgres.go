package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/ExtraProjects860/Project-Device-Mobile/handler/request"
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"github.com/jackc/pgx/v5/pgconn"
)

func verifyRoleDuplicated(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		if strings.Contains(pgErr.ConstraintName, "name") {
			return errors.New("this name role is registered")
		}
		return err
	}
	return err
}

func (r *PostgresRoleRepository) GetRole(ctx context.Context, id uint) (schemas.Role, error) {
	query := r.db.WithContext(ctx).Model(&schemas.Role{})

	role, err := getByID[schemas.Role](query, id)
	if err != nil {
		return schemas.Role{}, err
	}
	return role, nil
}

func (r *PostgresRoleRepository) CreateRole(ctx context.Context, role *schemas.Role) error {
	err := create(ctx, r.db, role)
	if err != nil {
		return verifyRoleDuplicated(err)
	}

	ro, err := r.GetRole(ctx, role.ID)
	if err != nil {
		return err
	}
	*role = ro

	return nil
}

func (r *PostgresRoleRepository) UpdateRole(ctx context.Context, id uint, role *schemas.Role) error {
	if err := updateByID(ctx, r.db, role, id); err != nil {
		return verifyRoleDuplicated(err)
	}

	ro, err := r.GetRole(ctx, role.ID)
	if err != nil {
		return err
	}
	*role = ro

	return nil
}

func (r *PostgresRoleRepository) GetRoles(ctx context.Context, paginationSearch request.PaginationSearch) ([]schemas.Role, uint, uint, error) {
	query := r.db.WithContext(ctx).Model(&schemas.Role{})

	roles, totalPages, totalItems, err := getByPagination[schemas.Role](
		query,
		paginationSearch,
	)
	if err != nil {
		return nil, 0, 0, err
	}

	return roles, totalPages, totalItems, err
}
