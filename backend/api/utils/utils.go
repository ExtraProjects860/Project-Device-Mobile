package utils

import (
	"math/rand"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GetValidate() *validator.Validate {
	return validate
}

func GenerateRandomCode(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = byte(rand.Intn(10) + '0')
	}
	return string(b)
}
