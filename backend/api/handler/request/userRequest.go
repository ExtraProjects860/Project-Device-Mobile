package request

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type UserCreateRequest struct {
	RoleID         uint    `json:"role_id" binding:"required"`
	EnterpriseID   *uint   `json:"enterprise_id"`
	Name           string  `json:"name" binding:"required"`
	Email          string  `json:"email" binding:"required,email"`
	Password       string  `json:"password" binding:"required,min=6"`
	Cpf            string  `json:"cpf" binding:"required"`
	RegisterNumber uint    `json:"register_number" binding:"required"`
	PhotoUrl       *string `json:"photo_url"`
}

// Pra que esse s de parametro? Tava já em outro lugar movi pra cá e ficou s de sexo

func (s *UserCreateRequest) Validate(ctx context.Context, validate *validator.Validate) error {
	if err := validate.Var(s.Name, "required,min=3"); err != nil {
		return fmt.Errorf("name: %v", err)
	}

	if err := validate.Var(s.Email, "required,email"); err != nil {
		return fmt.Errorf("email: %v", err)
	}

	if err := validate.Var(s.Password, "required,min=6"); err != nil {
		return fmt.Errorf("password: %v", err)
	}

	if err := validate.Var(s.Cpf, "required,cpf"); err != nil {
		return fmt.Errorf("cpf: %v", err)
	}

	if err := validate.Var(s.RegisterNumber, "required,gt=0"); err != nil {
		return fmt.Errorf("register_number: %v", err)
	}

	if err := validate.Var(s.RoleID, "required,gt=0"); err != nil {
		return fmt.Errorf("role_id: %v", err)
	}

	return nil
}

func (s *UserCreateRequest) ValidateUpdate(validate *validator.Validate) error {
	if s.Name != "" {
		if err := validate.Var(s.Name, "min=3"); err != nil {
			return fmt.Errorf("name: %v", err)
		}
	}

	if s.Email != "" {
		if err := validate.Var(s.Email, "email"); err != nil {
			return fmt.Errorf("email: %v", err)
		}
	}

	if s.Password != "" {
		if err := validate.Var(s.Password, "min=6"); err != nil {
			return fmt.Errorf("password: %v", err)
		}
	}

	if s.Cpf != "" {
		if err := validate.Var(s.Cpf, "len=11,numeric,cpf"); err != nil {
			return fmt.Errorf("cpf: %v", err)
		}
	}

	if s.RegisterNumber != 0 {
		if err := validate.Var(s.RegisterNumber, "gt=0"); err != nil {
			return fmt.Errorf("register_number: %v", err)
		}
	}

	if s.RoleID != 0 {
		if err := validate.Var(s.RoleID, "gt=0"); err != nil {
			return fmt.Errorf("role_id: %v", err)
		}
	}

	if s.EnterpriseID != nil && *s.EnterpriseID != 0 {
		if err := validate.Var(*s.EnterpriseID, "gt=0"); err != nil {
			return fmt.Errorf("enterprise_id: %v", err)
		}
	}

	return nil
}

func (s *UserCreateRequest) Format() {
	if s.Name != "" {
		s.Name = strings.ToUpper(strings.TrimSpace(s.Name))
	}

	if s.Email != "" {
		s.Email = strings.ToLower(strings.TrimSpace(s.Email))
	}

	if s.Cpf != "" {
		s.Cpf = strings.ReplaceAll(s.Cpf, ".", "")
		s.Cpf = strings.ReplaceAll(s.Cpf, "-", "")
		s.Cpf = strings.TrimSpace(s.Cpf)
	}

	if s.Password != "" {
		s.Password = strings.TrimSpace(s.Password)
	}

	if s.PhotoUrl != nil {
		photo := strings.TrimSpace(*s.PhotoUrl)
		s.PhotoUrl = &photo
	}
}
