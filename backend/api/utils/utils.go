package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func GenerateRandomPhoto(name string, width, height uint) string {
	return fmt.Sprintf(
		"https://picsum.photos/%d/%d?random=%s",
		width,
		height,
		name,
	)
}
