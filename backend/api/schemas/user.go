package schemas

import (
	"context"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	RoleID         uint `gorm:"not null"`
	EnterpriseID   *uint
	Name           string `gorm:"not null" validate:"required,min=3"`
	Email          string `gorm:"uniqueIndex;not null" validate:"required,email"`
	Password       string `gorm:"not null" validate:"required,min=10"`
	Cpf            string `gorm:"uniqueIndex;not null" validate:"required,cpf"`
	RegisterNumber uint   `gorm:"not null" validate:"required,min=3"`
	PhotoUrl       *string

	Role            Role          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Enterprise      Enterprise    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TokenPassword   TokenPassword `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WishListEntries []WishList    `gorm:"foreignKey:UserID"`
}

func (s *User) Validate(ctx context.Context, validate *validator.Validate) error {
	return validate.StructCtx(ctx, s)
}

func (s *User) Format() {
	s.Name = strings.ToUpper(strings.TrimSpace(s.Name))
	s.Email = strings.ToLower(strings.TrimSpace(s.Email))

	if s.PhotoUrl != nil {
		photo := strings.TrimSpace(*s.PhotoUrl)
		s.PhotoUrl = &photo
	}
}
