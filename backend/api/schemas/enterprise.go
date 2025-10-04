package schemas

import (
	"context"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Enterprise struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;not null" validate:"required,min=3"`
}

func (s *Enterprise) Validate(ctx context.Context, validate *validator.Validate) error {
	return validate.StructCtx(ctx, s)
}

func (s *Enterprise) Format() {
	s.Name = strings.ToUpper(strings.TrimSpace(s.Name))
}
