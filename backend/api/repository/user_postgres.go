package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

func verifyDuplicated(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		if strings.Contains(pgErr.ConstraintName, "email") {
			return errors.New("this email is registered")
		}
		if strings.Contains(pgErr.ConstraintName, "cpf") {
			return errors.New("this cpf is registered")
		}
		return err
	}
	return err
}

func (r *PostgresUserRepository) GetInfoUser(ctx context.Context, id uint) (schemas.User, error) {
	query := r.db.WithContext(ctx).Model(&schemas.User{}).Preload("Role").Preload("Enterprise")

	user, err := getByID[schemas.User](query, id)
	if err != nil {
		return schemas.User{}, err
	}
	return user, nil
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user *schemas.User) error {
	err := create(ctx, r.db, user)
	if err != nil {
		return verifyDuplicated(err)
	}

	u, err := r.GetInfoUser(ctx, user.ID)
	if err != nil {
		return err
	}
	*user = u

	return nil
}

func (r *PostgresUserRepository) UpdateUser(ctx context.Context, id uint, user *schemas.User) error {
	err := r.db.
		WithContext(ctx).
		Model(&schemas.User{}).
		Transaction(func(tx *gorm.DB) error {
			return tx.
				Where("id = ?", id).
				Updates(&user).Error
		})
	if err != nil {
		return verifyDuplicated(err)
	}

	u, err := r.GetInfoUser(ctx, user.ID)
	if err != nil {
		return err
	}
	*user = u

	return nil
}

func (r *PostgresUserRepository) GetUsers(ctx context.Context, itemsPerPage uint, currentPage uint) ([]schemas.User, uint, uint, error) {
	query := r.db.WithContext(ctx).Model(&schemas.User{}).Preload("Role").Preload("Enterprise")

	users, totalPages, totalItems, err := getByPagination[schemas.User](query, itemsPerPage, currentPage)
	if err != nil {
		return nil, 0, 0, err
	}

	return users, totalPages, totalItems, nil
}
