package dto

import (
	"time"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

type RoleDTO struct {
	ID        uint      `json:"id" example:"1"`
	Name      string    `json:"name" example:"admin"`
	CreatedAt time.Time `json:"created_at" example:"2025-10-12T21:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2025-10-12T21:05:00Z"`
}

func MakeRoleOutput(role schemas.Role) *RoleDTO {
	return &RoleDTO{
		ID:        role.ID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}
