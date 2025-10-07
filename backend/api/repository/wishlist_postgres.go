package repository

import (
	"context"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

func (r *postgresWishListRepository) GetWishListByUserID(ctx context.Context, userID uint, itemsPerPage uint, currentPage uint) ([]schemas.WishList, uint, uint, error) {
	query := r.db.WithContext(ctx).Where("user_id = ?", userID).Model(&schemas.WishList{}).Preload("Product")

	wishListEntries, totalPages, totalItems, err := getByPagination[schemas.WishList](query, itemsPerPage, currentPage)
	if err != nil {
		logger.Errorf("%v", err)
		return nil, 0, 0, err
	}

	return wishListEntries, totalPages, totalItems, err
}
