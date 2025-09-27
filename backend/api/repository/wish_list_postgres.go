package repository

import (
	"context"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

type WishListDTO struct {
	UserID    uint         `json:"id"`
	ItemCount int          `json:"item_count"`
	Products  []ProductDTO `json:"products"`
}

func makeWishListOutput(wishListEntries []schemas.WishList, userID uint) *WishListDTO {
	productsDTO := make([]ProductDTO, 0, len(wishListEntries))
	for _, entry := range wishListEntries {
		if entry.Product.ID == 0 {
			continue
		}

		product := ProductDTO{
			ID:          entry.Product.ID,
			Name:        entry.Product.Name,
			Description: entry.Product.Description,
			Value:       entry.Product.Value,
			PhotoUrl:    entry.Product.PhotoUrl,
		}
		productsDTO = append(productsDTO, product)
	}

	return &WishListDTO{
		UserID:    userID,
		ItemCount: len(productsDTO),
		Products:  productsDTO,
	}
}

func (r *postgresWishListRepository) GetWishListByUserID(ctx context.Context, userID uint, itemsPerPage uint, currentPage uint) (PaginationDTO, error) {
	query := r.db.WithContext(ctx).Where("user_id = ?", userID).Model(&schemas.WishList{})
	paginationOffset, totalPages := pagination(query, itemsPerPage, currentPage)

	var wishListEntries []schemas.WishList
	err := query.
		Limit(int(itemsPerPage)).
		Offset(int(paginationOffset)).
		Preload("Product").
		Find(&wishListEntries).Error
	if err != nil {
		logger.Errorf("%v", err)
		return PaginationDTO{}, err
	}

	wishListDTO := makeWishListOutput(wishListEntries, userID)

	return PaginationDTO{Data: wishListDTO, CurrentPage: currentPage, TotalPages: totalPages}, err
}
