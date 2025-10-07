package repository

import (
	"context"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

func (r *postgresTokenPasswordRepository) CreateToken(ctx context.Context, token schemas.TokenPassword) (schemas.TokenPassword, error) {
	err := create(ctx, r.db, &token)
	if err != nil {
		return schemas.TokenPassword{}, err
	}
	return token, nil
}

func (r *postgresTokenPasswordRepository) UpdateToken(ctx context.Context, id uint, token schemas.TokenPassword) (schemas.TokenPassword, error) {
	err := update(ctx, r.db, id, &token)
	if err != nil {
		return schemas.TokenPassword{}, err
	}
	return token, nil
}

func (r *postgresTokenPasswordRepository) GetToken(ctx context.Context, id uint) (schemas.TokenPassword, error) {
	query := r.db.WithContext(ctx).Model(&schemas.User{})

	token, err := getByID[schemas.TokenPassword](ctx, query, id)
	if err != nil {
		return schemas.TokenPassword{}, err
	}
	return token, nil
}
