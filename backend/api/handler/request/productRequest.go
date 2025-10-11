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
	IsPromotionAvaible *bool    `json:"is_promotion_avaible"`
	Discount           *float64 `json:"discount"`
	PhotoUrl           *string  `json:"photo_url"`
	IsAvaible          *bool    `json:"is_avaible"`
}

func (s *ProductRequest) Validate(validate *validator.Validate) error {
	if err := validate.Var(s.Name, "required,min=3"); err != nil {
		return fmt.Errorf("name: %v", err)
	}

	if err := validate.Var(s.Description, "required,min=3,max=255"); err != nil {
		return fmt.Errorf("description: %v", err)
	}

	if err := validate.Var(s.Value, "required,gt=0"); err != nil {
		return fmt.Errorf("value: %v", err)
	}

	if err := validate.Var(s.Quantity, "required,gte=0"); err != nil {
		return fmt.Errorf("quantity: %v", err)
	}

	if err := validate.Var(*s.Discount, "gte=0"); err != nil {
		return fmt.Errorf("discount: %v", err)
	}

	return nil
}

func (s *ProductRequest) ValidateUpdate() error {
	hasAtLeastOne := s.Name != "" ||
		s.Description != "" ||
		s.Value != 0 ||
		s.Quantity != 0 ||
		s.IsPromotionAvaible != nil ||
		*s.Discount != 0 && s.Discount != nil ||
		*s.PhotoUrl != "" && s.PhotoUrl != nil ||
		s.IsAvaible != nil

	if !hasAtLeastOne {
		return fmt.Errorf("at least one valid field must be provided")
	}

	return nil
}

func (s *ProductRequest) Format() {
	if s.Name != "" {
		name := strings.ToUpper(strings.TrimSpace(s.Name))
		s.Name = name
	}

	if s.Description != "" {
		description := strings.TrimSpace(s.Description)
		s.Description = description
	}

	if s.PhotoUrl != nil {
		photo := strings.TrimSpace(*s.PhotoUrl)
		s.PhotoUrl = &photo
	}
}
