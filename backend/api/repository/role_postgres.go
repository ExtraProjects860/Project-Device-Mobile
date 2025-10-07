package repository

import (
	"context"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

func (r *postgresRoleRepository) CreateRole(ctx context.Context, role schemas.Role) (schemas.Role, error) {
	err := create(ctx, r.db, &role)
	if err != nil {
		return schemas.Role{}, err
	}
	return role, nil
}

func (r *postgresRoleRepository) GetRoles(ctx context.Context, itemsPerPage uint, currentPage uint) ([]schemas.Role, uint, uint, error) {
	query := r.db.WithContext(ctx).Model(&schemas.Role{})

	roles, totalPages, totalItems, err := getByPagination[schemas.Role](query, itemsPerPage, currentPage)
	if err != nil {
		return nil, 0, 0, err
	}

	return roles, totalPages, totalItems, err
}

func (r *postgresRoleRepository) UpdateRole(ctx context.Context, id uint, role schemas.Role) (schemas.Role, error) {
	err := update(ctx, r.db, id, &role)
	if err != nil {
		return schemas.Role{}, err
	}
	return role, nil
}
