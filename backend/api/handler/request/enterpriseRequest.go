package request

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EnterpriseRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}

func (s *EnterpriseRequest) Validate(ctx *gin.Context, validate *validator.Validate) error {
	return validate.StructCtx(ctx, s)
}

func (s *EnterpriseRequest) ValidateUpdate() error {
	hasAtLeastOne := s.Name != ""
	if !hasAtLeastOne {
		return fmt.Errorf("at least one valid field must be provided")
	}
	return nil
}

func (s *EnterpriseRequest) Format() {
	if s.Name != "" {
		name := strings.ToUpper(strings.TrimSpace(s.Name))
		s.Name = name
	}
}
