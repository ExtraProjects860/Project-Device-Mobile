package request

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RoleRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}

func (s *RoleRequest) Validate(ctx *gin.Context, validate *validator.Validate) error {
	return validate.StructCtx(ctx, s)
}

func (s *RoleRequest) ValidateUpdate() error {
	hasAtLeastOne := s.Name != ""
	if !hasAtLeastOne {
		return fmt.Errorf("at least one valid field must be provided")
	}
	return nil
}

func (s *RoleRequest) Format() {
	if s.Name != "" {
		name := strings.ToUpper(strings.TrimSpace(s.Name))
		s.Name = name
	}
}
