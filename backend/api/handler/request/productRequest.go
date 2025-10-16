package request

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductRequest struct {
	Name               string   `json:"name" validate:"required,min=3"`
	Description        string   `json:"description" validate:"required,min=3,max=255"`
	Value              float64  `json:"value" validate:"required,gt=0"`
	Quantity           int      `json:"quantity" validate:"required,gte=0"`
	IsPromotionAvaible *bool    `json:"is_promotion_avaible"`
	Discount           *float64 `json:"discount" validate:"gte=0"`
	PhotoUrl           *string  `json:"photo_url"`
	IsAvaible          *bool    `json:"is_avaible"`
}

func (s *ProductRequest) Validate(ctx *gin.Context, validate *validator.Validate) error {
	return validate.StructCtx(ctx, s)
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
