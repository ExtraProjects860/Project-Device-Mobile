package dto

import (
	"time"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

type UserDTO struct {
	ID             uint      `json:"id" example:"1"`
	Name           string    `json:"name" example:"João da Silva"`
	Role           string    `json:"role" example:"admin"`
	Enterprise     *string   `json:"enterprise,omitempty" example:"Empresa XPTO"`
	Email          string    `json:"email" example:"joao@example.com"`
	Cpf            string    `json:"cpf" example:"123.456.789-00"`
	RegisterNumber string    `json:"register_number" example:"20210012"`
	PhotoUrl       *string   `json:"photo_url,omitempty" example:"https://cdn.exemplo.com/fotos/joao.jpg"`
	CreatedAt      time.Time `json:"created_at" example:"2025-10-12T21:00:00Z"`
	UpdatedAt      time.Time `json:"updated_at" example:"2025-10-12T21:05:00Z"`
}

func MakeUserOutput(user schemas.User) *UserDTO {
	var enterprise *string
	if user.Enterprise.ID != 0 {
		enterprise = &user.Enterprise.Name
	}

	return &UserDTO{
		ID:             user.ID,
		Name:           user.Name,
		Role:           user.Role.Name,
		Enterprise:     enterprise,
		Email:          user.Email,
		Cpf:            user.Cpf,
		RegisterNumber: user.RegisterNumber,
		PhotoUrl:       user.PhotoUrl,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
}
