package repository

import (
	"context"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

var logger *config.Logger = config.GetLogger("repository")

type EnterpriseRepository interface {
	CreateEnterprise(ctx context.Context, enterprise schemas.Enterprise) error
	GetEnterprises(ctx context.Context, itemsPerPage uint, currentPage uint) (PaginationDTO, error)
	UpdateEnterprise(ctx context.Context, id uint, enterprise schemas.Enterprise) (schemas.Enterprise, error)
}

type RoleRepository interface {
	CreateRole(ctx context.Context, role schemas.Role) error
	GetRoles(ctx context.Context, itemsPerPage uint, currentPage uint) (PaginationDTO, error)
	UpdateRole(ctx context.Context, id uint, role schemas.Role) (schemas.Role, error)
}

type WishListRepository interface {
	GetWishListByUserID(ctx context.Context, userID uint, itemsPerPage uint, currentPage uint) (PaginationDTO, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user schemas.User) error
	GetInfoUser(ctx context.Context, id uint) (*UserDTO, error)
	GetUsers(ctx context.Context, itemsPerPage uint, currentPage uint) (PaginationDTO, error)
	UpdateUser(ctx context.Context, id uint, user schemas.User) (schemas.User, error)
}

type TokenPasswordRepository interface {
	CreateToken(ctx context.Context, token schemas.TokenPassword) error
	UpdateToken(ctx context.Context, id uint, token schemas.TokenPassword) (schemas.TokenPassword, error)
	GetToken(ctx context.Context, id uint) (*TokenPasswordDTO, error)
}

type ProductRepository interface {
	CreateProduct(ctx context.Context, product schemas.Product) error
	GetProducts(ctx context.Context, itemsPerPage uint, currentPage uint) (PaginationDTO, error)
	UpdateProducts(ctx context.Context, id uint, product schemas.Product) (schemas.Product, error)
}
