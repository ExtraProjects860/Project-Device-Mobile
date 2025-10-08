package request

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type EnterpriseRequest struct {
	Name string `json:"name" binding:"required"`
}

func (e *EnterpriseRequest) Validate(validate *validator.Validate) error {
	if err := validate.Var(e.Name, "required,min=3"); err != nil {
		return fmt.Errorf("name: %v", err)
	}

	return nil
}

func (e *EnterpriseRequest) ValidateUpdate(validate *validator.Validate) error {
	if e.Name != "" {
		if err := validate.Var(e.Name, "min=3"); err != nil {
			return fmt.Errorf("name: %v", err)
		}
	}

	return nil
}

func (e *EnterpriseRequest) Format() {
	if e.Name != "" {
		e.Name = strings.ToUpper(strings.TrimSpace(e.Name))
	}
}
