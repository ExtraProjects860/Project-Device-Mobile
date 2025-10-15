package request

import (
	"fmt"
	"strings"

	"github.com/ExtraProjects860/Project-Device-Mobile/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserRequest struct {
	RoleID         uint    `json:"role_id" validate:"required,gt=0"`
	EnterpriseID   *uint   `json:"enterprise_id"`
	Name           string  `json:"name" validate:"required,min=3"`
	Email          string  `json:"email" validate:"required,email"`
	Password       string  `json:"password" validate:"required,min=6"`
	Cpf            string  `json:"cpf" validate:"required"`
	RegisterNumber uint    `json:"register_number" validate:"required,gt=0"`
	PhotoUrl       *string `json:"photo_url"`
}

func (s *UserRequest) Validate(ctx *gin.Context, validate *validator.Validate) error {
	if !utils.ValidateCPF(s.Cpf) {
		return fmt.Errorf("invalid cpf. Try other value")
	}

	if *s.EnterpriseID == 0 {
		return fmt.Errorf("enterprise can't zero. Try other value")
	}

	if s.PhotoUrl != nil && *s.PhotoUrl == "" {
		return fmt.Errorf("photo can't be empty")
	}

	return validate.StructCtx(ctx, s)
}

func (s *UserRequest) ValidateUpdate() error {
	hasAtLeastOne := s.Name != "" ||
		s.Email != "" ||
		s.Password != "" ||
		s.Cpf != "" ||
		s.RegisterNumber != 0 ||
		s.RoleID != 0 ||
		s.EnterpriseID != nil && *s.EnterpriseID != 0 ||
		s.PhotoUrl != nil && *s.PhotoUrl != ""

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
		cpf := utils.UnformatCPF(s.Cpf)
		s.Cpf = cpf
	}
	if s.Password != "" {
		password := strings.TrimSpace(s.Password)
		s.Password = password
	}
	if s.PhotoUrl != nil && *s.PhotoUrl != "" {
		photo := strings.TrimSpace(*s.PhotoUrl)
		s.PhotoUrl = &photo
	}
}
