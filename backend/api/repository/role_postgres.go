package repository

import (
	"context"
	"time"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

type RoleDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func makeRoleOutput(role schemas.Role) *RoleDTO {
	return &RoleDTO{
		ID:        role.ID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

func (r *postgresRoleRepository) CreateRole(ctx context.Context, role schemas.Role) error {
	return create(ctx, r.db, &role)
}

func (r *postgresRoleRepository) GetRoles(ctx context.Context, itemsPerPage uint, currentPage uint) (PaginationDTO, error) {
	query := r.db.WithContext(ctx).Model(&schemas.Role{})

	roles, totalPages, totalItems, err := getByPagination[schemas.Role](query, itemsPerPage, currentPage)
	if err != nil {
		return PaginationDTO{}, err
	}

	var rolesDTO []RoleDTO
	for _, role := range roles {
		rolesDTO = append(rolesDTO, *makeRoleOutput(role))
	}

	return PaginationDTO{
		Data:        rolesDTO,
		CurrentPage: currentPage,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}, nil
}

func (r *postgresRoleRepository) UpdateRole(ctx context.Context, id uint, role schemas.Role) (schemas.Role, error) {
	err := update(ctx, r.db, id, &role)
	return role, err
}
