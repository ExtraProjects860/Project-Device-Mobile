package repository

import "time"

type EnterprisePostgres interface {
	CreateEnterprise()
	GetEnterprises()
	UpdateEnterprise()
}

type EnterpriseDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
