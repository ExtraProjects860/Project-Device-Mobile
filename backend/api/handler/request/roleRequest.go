package request

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type RoleRequest struct {
	Name string `json:"name" binding:"required"`
}

func (r *RoleRequest) Validate(validate *validator.Validate) error {
	if err := validate.Var(r.Name, "required,min=3"); err != nil {
		return fmt.Errorf("name: %v", err)
	}

	return nil
}

func (r *RoleRequest) ValidateUpdate(validate *validator.Validate) error {
	if r.Name != "" {
		if err := validate.Var(r.Name, "min=3"); err != nil {
			return fmt.Errorf("name: %v", err)
		}
	}

	return nil
}

func (r *RoleRequest) Format() {
	if r.Name != "" {
		r.Name = strings.ToUpper(strings.TrimSpace(r.Name))
	}
}
