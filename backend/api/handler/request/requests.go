package request

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidateModel interface {
	Format()
	Validate(v *validator.Validate) error
	ValidateUpdate(v *validator.Validate) error
}

func ValidateBodyReq(v ValidateModel, val *validator.Validate) error {
	err := v.Validate(val)
	v.Format()
	return err
}

func ValidateUpdateBodyReq(v ValidateModel, val *validator.Validate) error {
	err := v.ValidateUpdate(val)
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

func GetPaginationData(ctx *gin.Context) (uint, uint, error) {
	itemsPerPage := ctx.Query("itemsPerPage")
	if itemsPerPage == "" {
		return 0, 0, ErrParamIsRequired("id", "queryParameter")
	}

	parsedItemsPerPage, err := strconv.ParseUint(itemsPerPage, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid parsed itemsPerPage: %v", err)
	}

	currentPage := ctx.Query("currentPage")
	if currentPage == "" {
		return 0, 0, ErrParamIsRequired("id", "queryParameter")
	}

	parsedCurrentPage, err := strconv.ParseUint(currentPage, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid parsed currentPage: %v", err)
	}

	return uint(parsedItemsPerPage), uint(parsedCurrentPage), nil
}
