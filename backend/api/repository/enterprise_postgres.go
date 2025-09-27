package repository

import (
	"context"
	"time"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

type EnterpriseDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func makeEnterpriseOutput(enterprise schemas.Enterprise) *EnterpriseDTO {
	return &EnterpriseDTO{
		ID:        enterprise.ID,
		Name:      enterprise.Name,
		CreatedAt: enterprise.CreatedAt,
		UpdatedAt: enterprise.UpdatedAt,
	}
}

func (r *postgresEnterpriseRepository) CreateEnterprise(ctx context.Context, enterprise schemas.Enterprise) {
	return
}

func (r *postgresEnterpriseRepository) GetEnterprises(ctx context.Context, id uint) {
	return
}

func (r *postgresEnterpriseRepository) UpdateEnterprise(ctx context.Context, id uint, enterprise schemas.Enterprise) {
	return
}
