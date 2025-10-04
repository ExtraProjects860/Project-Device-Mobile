package repository

import (
	"context"
	"time"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

type UserDTO struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Role           string    `json:"role"`
	Enterprise     *string   `json:"enterprise,omitempty"`
	Email          string    `json:"email"`
	Cpf            string    `json:"cpf"`
	RegisterNumber uint      `json:"register_number"`
	PhotoUrl       *string   `json:"photo_url,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func makeUserOutput(user schemas.User) *UserDTO {
	var enterprise *string
	if user.Enterprise.ID != 0 {
		enterprise = &user.Enterprise.Name
	}

	return &UserDTO{
		ID:             user.ID,
		Name:           user.Name,
		Role:           user.Role.Name,
		Enterprise:     enterprise,
		Email:          user.Email,
		Cpf:            user.Cpf,
		RegisterNumber: user.RegisterNumber,
		PhotoUrl:       user.PhotoUrl,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
}

func (r *postgresUserRepository) CreateUser(ctx context.Context, user schemas.User) error {
	return create(ctx, r.db, &user)
}

func (r *postgresUserRepository) GetInfoUser(ctx context.Context, id uint) (*UserDTO, error) {
	query := r.db.WithContext(ctx).Model(&schemas.User{}).Preload("Role").Preload("Enterprise")

	user, err := getByID[schemas.User](ctx, query, id)
	if err != nil {
		return nil, err
	}
	return makeUserOutput(*user), nil
}

func (r *postgresUserRepository) UpdateUser(ctx context.Context, id uint, user schemas.User) (schemas.User, error) {
	err := update(ctx, r.db, id, &user)
	return user, err
}

func (r *postgresUserRepository) GetUsers(ctx context.Context, itemsPerPage uint, currentPage uint) (PaginationDTO, error) {
	query := r.db.WithContext(ctx).Model(&schemas.User{}).Preload("Role").Preload("Enterprise")

	users, totalPages, totalItems, err := getByPagination[schemas.User](query, itemsPerPage, currentPage)
	if err != nil {
		return PaginationDTO{}, err
	}

	var usersDTO []UserDTO
	for _, user := range users {
		usersDTO = append(usersDTO, *makeUserOutput(user))
	}

	return PaginationDTO{
		Data:        usersDTO,
		CurrentPage: currentPage,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}, nil
}
