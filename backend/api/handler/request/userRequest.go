package request

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type UserRequest struct {
	RoleID         uint    `json:"role_id" `
	EnterpriseID   *uint   `json:"enterprise_id"`
	Name           string  `json:"name" `
	Email          string  `json:"email" `
	Password       string  `json:"password" `
	Cpf            string  `json:"cpf" `
	RegisterNumber uint    `json:"register_number" `
	PhotoUrl       *string `json:"photo_url"`
}

// Pra que esse s de parametro? Tava já em outro lugar movi pra cá e ficou s de sexo

func (s *UserRequest) Validate(validate *validator.Validate) error {
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

func (s *UserRequest) ValidateUpdate() error {
	hasAtLeastOne := s.Name != "" ||
		s.Email != "" ||
		s.Password != "" ||
		s.Cpf != "" ||
		s.RegisterNumber != 0 ||
		s.RoleID != 0 ||
		s.EnterpriseID != nil ||
		s.PhotoUrl != nil

	if !hasAtLeastOne {
		return fmt.Errorf("at least one valid field must be provided")
	}
	return nil
}

func (s *UserRequest) Format() {
	if s.Name != "" {
		name := strings.ToUpper(strings.TrimSpace(s.Name))
		s.Name = name
	}

	if s.Email != "" {
		email := strings.ToLower(strings.TrimSpace(s.Email))
		s.Email = email
	}

	if s.Cpf != "" {
		cpf := strings.ReplaceAll(s.Cpf, ".", "")
		cpf = strings.ReplaceAll(cpf, "-", "")
		cpf = strings.TrimSpace(cpf)
		s.Cpf = cpf
	}

	if s.Password != "" {
		password := strings.TrimSpace(s.Password)
		s.Password = password
	}

	if s.PhotoUrl != nil {
		photo := strings.TrimSpace(*s.PhotoUrl)
		s.PhotoUrl = &photo
	}
}
