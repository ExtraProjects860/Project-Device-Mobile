package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// JWT Token

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserData struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

func errParamIsRequired(name string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}
func getIdQuery(ctx *gin.Context) (uint, error) {
	id := ctx.Query("id")
	if id == "" {
		return 0, errParamIsRequired("id", "queryParameter")
	}

	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid id: %v", err)
	}

	return uint(parsedId), nil
}

func getPaginationData(ctx *gin.Context) (uint, uint, error) {
	itemsPerPage := ctx.Query("itemsPerPage")
	if itemsPerPage == "" {
		return 0, 0, errParamIsRequired("id", "queryParameter")
	}

	parsedItemsPerPage, err := strconv.ParseUint(itemsPerPage, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid parsed itemsPerPage: %v", err)
	}

	currentPage := ctx.Query("currentPage")
	if currentPage == "" {
		return 0, 0, errParamIsRequired("id", "queryParameter")
	}

	parsedCurrentPage, err := strconv.ParseUint(currentPage, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid parsed currentPage: %v", err)
	}

	return uint(parsedItemsPerPage), uint(parsedCurrentPage), nil
}
