package repository

import (
	"context"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

// TODO preciso construir os métodos aqui dps

func (r *PostgresAuthRepository) CreateToken(ctx context.Context, token *schemas.TokenPassword) (*schemas.TokenPassword, error) {
	return &schemas.TokenPassword{}, nil
}

func (r *PostgresAuthRepository) UpdateToken(ctx context.Context, id uint, token *schemas.TokenPassword) (schemas.TokenPassword, error) {
	panic("Not implemented")
}

func (r *PostgresAuthRepository) GetToken(ctx context.Context, id uint) (*schemas.TokenPassword, error) {
	return &schemas.TokenPassword{}, nil
}
