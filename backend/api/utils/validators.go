package utils

import (
	"github.com/go-playground/validator/v10"
)

func GetValidate() *validator.Validate {
	return validate
}
