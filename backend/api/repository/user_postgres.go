package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"github.com/jackc/pgx/v5/pgconn"
)

func verifyUserDuplicated(err error) error {
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

// TODO ajustar erro para dar notfound e não só "error to process get user"
func (r *PostgresUserRepository) GetInfoUser(ctx context.Context, id uint) (schemas.User, error) {
	query := r.db.WithContext(ctx).Model(&schemas.User{}).Preload("Role").Preload("Enterprise")

	user, err := getByID[schemas.User](query, id)
	if err != nil {
		return schemas.User{}, err
	}
	return user, nil
}

func (r *PostgresUserRepository) GetUserByEmail(ctx context.Context, email string) (schemas.User, error) {
	model := r.db.WithContext(ctx).Model(&schemas.User{})
	user, err := firstWhere[schemas.User](
		model,
		"email = ?",
		email,
	)
	if err != nil {
		return schemas.User{}, err
	}

	return user, nil
}


func (r *PostgresUserRepository) CreateUser(ctx context.Context, user *schemas.User) error {
	err := create(ctx, r.db, user)
	if err != nil {
		return verifyUserDuplicated(err)
	}

	u, err := r.GetInfoUser(ctx, user.ID)
	if err != nil {
		return err
	}
	*user = u

	return nil
}

// TODO na hora de atualizar a senha é só meter o update, burro da 0 pra ele

func (r *PostgresUserRepository) UpdateUser(ctx context.Context, id uint, user *schemas.User) error {
	if err := updateByID(ctx, r.db, user, id); err != nil {
		return verifyUserDuplicated(err)
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

	users, totalPages, totalItems, err := getByPagination[schemas.User](
		query,
		itemsPerPage,
		currentPage,
	)
	if err != nil {
		return nil, 0, 0, err
	}

	return users, totalPages, totalItems, nil
}
