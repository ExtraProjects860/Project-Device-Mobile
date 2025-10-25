package request

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidateModel interface {
	Format()
	Validate(ctx *gin.Context, val *validator.Validate) error
	ValidateUpdate() error
}

func ValidateBodyReq(v ValidateModel, ctx *gin.Context, val *validator.Validate) error {
	err := v.Validate(ctx, val)
	v.Format()
	return err
}

func ValidateUpdateBodyReq(v ValidateModel) error {
	err := v.ValidateUpdate()
	v.Format()
	return err
}

func ErrParamIsRequired(name string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func GetIdQuery(ctx *gin.Context) (uint, error) {
	id := ctx.Query("id")
	if id == "" {
		return 0, ErrParamIsRequired("id", "queryParameter")
	}

	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid id: %v", err)
	}

	return uint(parsedId), nil
}

func ReadBody[T any](ctx *gin.Context, input *T) error {
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return fmt.Errorf("error to parsed body to json")
	}

	return nil
}
