package request

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoginRequest struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
	RememberMe *bool  `json:"remember_me"`
}

type ChangePassword struct {
	NewPassword string `json:"new_password" validate:"required"`
}

func (s *LoginRequest) Validate(ctx *gin.Context, validate *validator.Validate) error {
	return validate.StructCtx(ctx, s)
}

func (s *LoginRequest) ValidateUpdate() error {
	return nil
}

func (s *LoginRequest) Format() {
	if s.Email != "" {
		email := strings.ToLower(strings.TrimSpace(s.Email))
		s.Email = email
	}
	if s.Password != "" {
		password := strings.TrimSpace(s.Password)
		s.Password = password
	}
	if s.RememberMe != nil {
		rememberMe := *s.RememberMe || false
		s.RememberMe = &rememberMe
	}
}
