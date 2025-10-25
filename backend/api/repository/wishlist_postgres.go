package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/ExtraProjects860/Project-Device-Mobile/handler/request"
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"github.com/jackc/pgx/v5/pgconn"
)

func verifyWishListDuplicated(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		if strings.Contains(pgErr.ConstraintName, "fk_products_wish_list_entries") ||
			strings.Contains(pgErr.ConstraintName, "fk_users_wish_list_entries") {
			return errors.New("this product is already in the user's wishlist")
		}
		return err
	}
	return err
}

func (r *PostgresWishListRepository) AddInWishList(ctx context.Context, wishlist *schemas.WishList) error {
	err := create(ctx, r.db, wishlist)
	if err != nil {
		return verifyWishListDuplicated(err)
	}

	return nil
}

func (r *PostgresWishListRepository) DeleteInWishList(ctx context.Context, userID, productID uint) error {
	return delete(
		ctx,
		r.db,
		&schemas.WishList{},
		"user_id = ? AND product_id = ?",
		userID, productID,
	)
}

func (r *PostgresWishListRepository) GetWishListByUserID(ctx context.Context, userID uint, paginationSearch request.PaginationSearch) ([]schemas.WishList, uint, uint, error) {
	query := r.db.WithContext(ctx).Where("user_id = ?", userID).Model(&schemas.WishList{}).Preload("Product")

	wishListEntries, totalPages, totalItems, err := getByPagination[schemas.WishList](
		query,
		paginationSearch,
	)
	if err != nil {
		return nil, 0, 0, err
	}

	return wishListEntries, totalPages, totalItems, err
}
