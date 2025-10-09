package request

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ProductRequest struct {
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	Value              float64  `json:"value"`
	Quantity           int      `json:"quantity"`
	IsPromotionAvaible bool     `json:"is_promotion_avaible"`
	Discount           *float64 `json:"discount"`
	PhotoUrl           *string  `json:"photo_url"`
	IsAvaible          bool     `json:"is_avaible"`
}

func (p *ProductRequest) Validate(validate *validator.Validate) error {
	if err := validate.Var(p.Name, "required,min=3"); err != nil {
		return fmt.Errorf("name: %v", err)
	}

	if err := validate.Var(p.Description, "required,min=3,max=255"); err != nil {
		return fmt.Errorf("description: %v", err)
	}

	if err := validate.Var(p.Value, "required,gt=0"); err != nil {
		return fmt.Errorf("value: %v", err)
	}

	if err := validate.Var(p.Quantity, "required,gte=0"); err != nil {
		return fmt.Errorf("quantity: %v", err)
	}

	if p.Discount != nil {
		if err := validate.Var(*p.Discount, "gte=0"); err != nil {
			return fmt.Errorf("discount: %v", err)
		}
	}

	return nil
}

func (p *ProductRequest) ValidateUpdate(validate *validator.Validate) error {
	if p.Name != "" {
		if err := validate.Var(p.Name, "min=3"); err != nil {
			return fmt.Errorf("name: %v", err)
		}
	}

	if p.Description != "" {
		if err := validate.Var(p.Description, "min=3,max=255"); err != nil {
			return fmt.Errorf("description: %v", err)
		}
	}

	if p.Value != 0 {
		if err := validate.Var(p.Value, "gt=0"); err != nil {
			return fmt.Errorf("value: %v", err)
		}
	}

	if p.Quantity != 0 {
		if err := validate.Var(p.Quantity, "gte=0"); err != nil {
			return fmt.Errorf("quantity: %v", err)
		}
	}

	if p.Discount != nil && *p.Discount != 0 {
		if err := validate.Var(*p.Discount, "gte=0"); err != nil {
			return fmt.Errorf("discount: %v", err)
		}
	}

	return nil
}

func (p *ProductRequest) Format() {
	if p.Name != "" {
		p.Name = strings.ToUpper(strings.TrimSpace(p.Name))
	}

	if p.Description != "" {
		p.Description = strings.TrimSpace(p.Description)
	}

	if p.PhotoUrl != nil {
		photo := strings.TrimSpace(*p.PhotoUrl)
		p.PhotoUrl = &photo
	}
}
