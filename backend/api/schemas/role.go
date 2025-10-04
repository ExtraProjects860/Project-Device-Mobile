package schemas

import (
	"context"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name  string `gorm:"unique;not null" validate:"required,min=3"`
	Users []User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (s *Role) Validate(ctx context.Context, validate *validator.Validate) error {
	return validate.StructCtx(ctx, s)
}

func (s *Role) Format() {
	s.Name = strings.ToUpper(strings.TrimSpace(s.Name))
}
