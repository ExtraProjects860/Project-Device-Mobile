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

func makeRoleDTO(role schemas.Role) *RoleDTO {
	return &RoleDTO{
		ID:        role.ID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

func (r *postgresRoleRepository) CreateRole(ctx context.Context, role schemas.Role) {
	return
}

func (r *postgresRoleRepository) GetRoles(ctx context.Context, id uint) {
	return
}

func (r *postgresRoleRepository) UpdateRole(ctx context.Context, id uint, role schemas.Role) {
	return
}
