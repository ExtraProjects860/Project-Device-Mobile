package handler

import (
	"errors"
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/repository"
	"github.com/gin-gonic/gin"
)

type WishListHandler struct {
	repo repository.WishListRepository
}

func NewWishListHandler(repo repository.WishListRepository) *WishListHandler {
	return &WishListHandler{repo: repo}
}

// @Summary      Add Product to Wish List
// @Description  Adds a product to the user wish list
// @Tags         wishlist
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /api/v1/wishlist [post]
func (h *WishListHandler) AddInWishListHandler(ctx *gin.Context) {
	sendSuccess(ctx, http.StatusCreated, "Add Product in Wish List!")
}

// @Summary      Update Product from Wish List
// @Description  Update a product from the user wish list
// @Tags         wishlist
// @Param 		 id query string true "WishList ID"
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /api/v1/wishlist [patch]
func (h *WishListHandler) UpdateWishListHandler(ctx *gin.Context) {
	sendSuccess(ctx, http.StatusOK, "Delete Product in Wish List!")
}

// @Summary      Get Wish List Items
// @Description  Returns all products in the user wish list
// @Tags         wishlist
// @Produce      json
// @Param 		 id query string true "User ID"
// @Param        itemsPerPage query string true "Pagination Items"
// @Param        currentPage query string true "Pagination Current Page"
// @Success      200 {array}  repository.WishListDTO
// @Failure      400 {object} ErrResponse
// @Failure      500 {object} ErrResponse
// @Router       /api/v1/wishlist [get]
func (h *WishListHandler) GetWishListByUserIDHandler(ctx *gin.Context) {
	userId, err := getIdQuery(ctx)
	if err != nil {
		logger.Error(err.Error())
		sendErr(ctx, http.StatusBadRequest, err)
		return
	}

	itemsPerPage, currentPage, err := getPaginationData(ctx)
	if err != nil {
		logger.Error(err.Error())
		sendErr(ctx, http.StatusBadRequest, err)
		return
	}

	wishlist, err := h.repo.GetWishListByUserID(ctx, userId, itemsPerPage, currentPage)
	if err != nil {
		logger.Error(err.Error())
		sendErr(ctx, http.StatusInternalServerError, errors.New("error to get wishlist in database"))
		return
	}

	sendSuccess(ctx, http.StatusOK, wishlist)
}
