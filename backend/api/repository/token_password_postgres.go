package repository

import (
	"context"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

func (r *PostgresTokenPasswordRepository) CreateToken(ctx context.Context, token schemas.TokenPassword) (schemas.TokenPassword, error) {
	err := create(ctx, r.db, &token)
	if err != nil {
		return schemas.TokenPassword{}, err
	}
	return token, nil
}

func (r *PostgresTokenPasswordRepository) UpdateToken(ctx context.Context, id uint, token schemas.TokenPassword) (schemas.TokenPassword, error) {
	panic("Not implemented")
}

func (r *PostgresTokenPasswordRepository) GetToken(ctx context.Context, id uint) (schemas.TokenPassword, error) {
	query := r.db.WithContext(ctx).Model(&schemas.User{})

	token, err := getByID[schemas.TokenPassword](query, id)
	if err != nil {
		return schemas.TokenPassword{}, err
	}
	return token, nil
}
