package repository

import (
	"context"
	"time"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user schemas.User) error
	GetInfoUser(id uint) (schemas.User, error)
	GetUsers() ([]schemas.User, error)
	UpdateUser(id uint) (schemas.User, error)
}

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

// TODO colocar mensagens nos logs aqui depois

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user schemas.User) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			logger.Errorf("%v", err)
			return err
		}
		return nil
	})
}

func (r *PostgresUserRepository) GetInfoUser(ctx context.Context, id uint) (*UserDTO, error) {
	var user schemas.User
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Preload("Role").
			Preload("Enterprise").
			First(&user, id).Error; err != nil {
			logger.Errorf("%v", err)
			return err
		}

		return nil
	})

	return makeUserOutput(user), err
}

func (r *PostgresUserRepository) GetUsers(ctx context.Context) ([]UserDTO, error) {
	var users []schemas.User
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Preload("Role").
			Preload("Enterprise").
			Find(&users).Error; err != nil {
			logger.Errorf("%v", err)
			return err
		}

		return nil
	})

	var usersDTO []UserDTO
	for _, user := range users {
		usersDTO = append(usersDTO, *makeUserOutput(user))
	}

	return usersDTO, err
}

func (r *PostgresUserRepository) UpdateUser(ctx context.Context, id uint, user schemas.User) (schemas.User, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&schemas.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
			logger.Errorf("%v", err)
			return err
		}

		return nil
	})

	return user, err
}
