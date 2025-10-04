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

func (r *postgresEnterpriseRepository) CreateEnterprise(ctx context.Context, enterprise schemas.Enterprise) error {
	return create(ctx, r.db, &enterprise)
}

func (r *postgresEnterpriseRepository) GetEnterprises(ctx context.Context, itemsPerPage uint, currentPage uint) (PaginationDTO, error) {
	query := r.db.WithContext(ctx).Model(&schemas.Enterprise{})

	enterprises, totalPages, totalItems, err := getByPagination[schemas.Enterprise](query, itemsPerPage, currentPage)
	if err != nil {
		return PaginationDTO{}, err
	}

	var enterprisesDTO []EnterpriseDTO
	for _, enterprise := range enterprises {
		enterprisesDTO = append(enterprisesDTO, *makeEnterpriseOutput(enterprise))
	}

	return PaginationDTO{
		Data:        enterprisesDTO,
		CurrentPage: currentPage,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}, nil
}

func (r *postgresEnterpriseRepository) UpdateEnterprise(ctx context.Context, id uint, enterprise schemas.Enterprise) (schemas.Enterprise, error) {
	err := update(ctx, r.db, id, &enterprise)
	return enterprise, err
}
