package dto

import (
	"time"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

type EnterpriseDTO struct {
	ID        uint      `json:"id" example:"1"`
	Name      string    `json:"name" example:"Empresa XPTO"`
	CreatedAt time.Time `json:"created_at" example:"2025-10-10T08:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2025-10-12T09:30:00Z"`
}

func MakeEnterpriseOutput(enterprise schemas.Enterprise) *EnterpriseDTO {
	return &EnterpriseDTO{
		ID:        enterprise.ID,
		Name:      enterprise.Name,
		CreatedAt: enterprise.CreatedAt,
		UpdatedAt: enterprise.UpdatedAt,
	}
}
