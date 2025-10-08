package dto

import (
	"time"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

type EnterpriseDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MakeEnterpriseOutput(enterprise schemas.Enterprise) *EnterpriseDTO {
	return &EnterpriseDTO{
		ID:        enterprise.ID,
		Name:      enterprise.Name,
		CreatedAt: enterprise.CreatedAt,
		UpdatedAt: enterprise.UpdatedAt,
	}
}
