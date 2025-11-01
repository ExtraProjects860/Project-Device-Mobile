package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

/*
TODO essa parte do código merece atenção para ser melhorada em alguns pontos
*/

func verifyTokenDuplicated(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		if strings.Contains(pgErr.ConstraintName, "user_id") {
			return errors.New("this user_id is registered")
		}
		return err
	}
	return err
}

func (r *PostgresAuthRepository) CreateToken(ctx context.Context, token *schemas.TokenPassword) error {
	err := create(ctx, r.db, token)
	if err != nil {
		return err
	}

	t, err := getByID[schemas.TokenPassword](r.db, token.ID)
	if err != nil {
		return err
	}
	*token = t

	return nil
}

func (r *PostgresAuthRepository) UpdateToken(ctx context.Context, id uint, token *schemas.TokenPassword) error {
	if err := updateByID(ctx, r.db, token, id); err != nil {
		return verifyTokenDuplicated(err)
	}

	t, err := r.GetToken(ctx, token.ID)
	if err != nil {
		return err
	}
	*token = t

	return nil
}

func (r *PostgresAuthRepository) GetToken(ctx context.Context, id uint) (schemas.TokenPassword, error) {
	token, err := firstWhere[schemas.TokenPassword](
		r.db,
		"user_id = ?", id,
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return schemas.TokenPassword{}, nil
		}
		return schemas.TokenPassword{}, err
	}

	return token, nil
}
