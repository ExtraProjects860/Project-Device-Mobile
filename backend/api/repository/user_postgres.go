package repository

import (
	"context"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)



func (r *postgresUserRepository) CreateUser(ctx context.Context, user schemas.User) (schemas.User, error) {
	err := create(ctx, r.db, &user)
	if err != nil {
		return schemas.User{}, err
	}
	return user, nil
}

func (r *postgresUserRepository) GetInfoUser(ctx context.Context, id uint) (schemas.User, error) {
	query := r.db.WithContext(ctx).Model(&schemas.User{}).Preload("Role").Preload("Enterprise")

	user, err := getByID[schemas.User](ctx, query, id)
	if err != nil {
		return schemas.User{}, err
	}
	return user, nil
}

func (r *postgresUserRepository) UpdateUser(ctx context.Context, id uint, user schemas.User) (schemas.User, error) {
	err := update(ctx, r.db, id, &user)

	if err != nil {
		return schemas.User{}, err
	}
	return user, nil
}

func (r *postgresUserRepository) GetUsers(ctx context.Context, itemsPerPage uint, currentPage uint) ([]schemas.User, uint, uint, error) {
	query := r.db.WithContext(ctx).Model(&schemas.User{}).Preload("Role").Preload("Enterprise")

	users, totalPages, totalItems, err := getByPagination[schemas.User](query, itemsPerPage, currentPage)
	if err != nil {
		return nil, 0, 0, err
	}

	return users, totalPages, totalItems, nil
}
