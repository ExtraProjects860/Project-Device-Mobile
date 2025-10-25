package request

import (
	"fmt"
	"strconv"

	"github.com/ExtraProjects860/Project-Device-Mobile/enum"
	"github.com/gin-gonic/gin"
)

type PaginationSearch struct {
	ItemsPerPage uint
	CurrentPage  uint
	SearchFilter string
	ItemsOrder   enum.ItemsOrder
}

func parsePages(itemsPerPage, currentPage string) (uint, uint, error) {
	parsedCurrentPage, err := strconv.ParseUint(currentPage, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid parsed currentPage: %v", err)
	}

	parsedItemsPerPage, err := strconv.ParseUint(itemsPerPage, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid parsed itemsPerPage: %v", err)
	}

	return uint(parsedCurrentPage), uint(parsedItemsPerPage), nil
}

func GetPaginationData(ctx *gin.Context) (PaginationSearch, error) {
	itemsPerPage := ctx.Query("itemsPerPage")
	if itemsPerPage == "" {
		return PaginationSearch{}, ErrParamIsRequired("id", "queryParameter")
	}

	currentPage := ctx.Query("currentPage")
	if currentPage == "" {
		return PaginationSearch{}, ErrParamIsRequired("id", "queryParameter")
	}

	parsedCurrentPage, parsedItemsPerPage, err := parsePages(itemsPerPage, currentPage)
	if err != nil {
		return PaginationSearch{}, err
	}

	searchFilter := ctx.Query("searchFilter")
	itemsOrderStr := ctx.Query("itemsOrder")

	itemsOrder, err := enum.ParseItemsOrder(itemsOrderStr)
	if err != nil {
		itemsOrder = enum.ASC
	}

	return PaginationSearch{
		ItemsPerPage: parsedItemsPerPage,
		CurrentPage:  parsedCurrentPage,
		SearchFilter: searchFilter,
		ItemsOrder:   itemsOrder,
	}, nil
}
