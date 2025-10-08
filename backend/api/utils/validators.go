package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/paemuri/brdoc"
)

func validateCPF(fl validator.FieldLevel) bool {
	cpf := fl.Field().String()
	return brdoc.IsCPF(cpf)
}

func GetValidate() *validator.Validate {
	return validate
}
