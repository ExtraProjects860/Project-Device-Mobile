package utils

import (
	"context"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"github.com/go-playground/validator/v10"
	"github.com/paemuri/brdoc"
)

func validateCPF(fl validator.FieldLevel) bool {
	cpf := fl.Field().String()
	return brdoc.IsCPF(cpf)
}

func PrepareAndValidate(ctx context.Context, schema schemas.ValidateModel, validate *validator.Validate) error {
	schema.Format()
	return schema.Validate(ctx, validate)
}

func GetValidate() *validator.Validate {
	return validate
}
