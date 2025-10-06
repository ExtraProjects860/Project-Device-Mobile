package dto

import (
	"time"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

type UserDTO struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Role           string    `json:"role"`
	Enterprise     *string   `json:"enterprise,omitempty"`
	Email          string    `json:"email"`
	Cpf            string    `json:"cpf"`
	RegisterNumber uint      `json:"register_number"`
	PhotoUrl       *string   `json:"photo_url,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
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
