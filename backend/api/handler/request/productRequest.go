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

// Arquivo: /app/handler/request/productRequest.go

func (s *ProductRequest) ValidateUpdate() error {
    if s.Name != "" { return nil }
    if s.Description != "" { return nil }
    if s.Value != 0 { return nil }
    if s.Quantity != 0 { return nil }
    if s.IsPromotionAvaible != nil { return nil }
    if s.Discount != nil { return nil }
    if s.PhotoUrl != nil { return nil }
    if s.IsAvaible != nil { return nil }
    return fmt.Errorf("at least one field must be provided for update")
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
