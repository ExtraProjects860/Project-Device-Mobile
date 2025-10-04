package schemas

import (
	"context"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name               string   `gorm:"uniqueIndex;not null" validate:"required,min=3"`
	Description        string   `gorm:"not null" validate:"required,min=3,max=255"`
	Value              float64  `gorm:"type:decimal(10,2);not null" validate:"required,gt=0"`
	Quantity           int      `gorm:"not null" validate:"required,gte=0"`
	IsPromotionAvaible bool     `gorm:"not null"`
	Discount           *float64 `gorm:"type:decimal(10,2)" validate:"omitempty,gte=0"`
	PhotoUrl           *string
	IsAvaible          bool `gorm:"default:true;not null"`

	WishListEntries []WishList `gorm:"foreignKey:ProductID"`
}

func (s *Product) Validate(ctx context.Context, validate *validator.Validate) error {
	return validate.StructCtx(ctx, s)
}

func (s *Product) Format() {
	s.Name = strings.ToUpper(strings.TrimSpace(s.Name))
	s.Description = strings.TrimSpace(s.Description)

	if s.PhotoUrl != nil {
		photo := strings.TrimSpace(*s.PhotoUrl)
		s.PhotoUrl = &photo
	}
}
