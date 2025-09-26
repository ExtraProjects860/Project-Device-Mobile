package repository

import (
	"context"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeRepository() {
	db = config.GetDB()
	logger = config.GetLogger("repository")
}

type EnterpriseRepository interface {
	CreateEnterprise(ctx context.Context, enterprise schemas.Enterprise)
	GetEnterprises(ctx context.Context, id uint)
	UpdateEnterprise(ctx context.Context, id uint, enterprise schemas.Enterprise)
}

type WishListRepository interface {
	AddInWishList(ctx context.Context, wishlist schemas.WishList)
	GetItemsWishList(ctx context.Context, itemsPerPage uint, currentPage uint)
	UpdateWishList(ctx context.Context, id uint, wishlist schemas.WishList)
}

type RoleRepository interface {
	CreateRole(ctx context.Context, role schemas.Role)
	GetRoles(ctx context.Context, id uint)
	UpdateRole(ctx context.Context, id uint, role schemas.Role)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user schemas.User) error
	GetInfoUser(ctx context.Context, id uint) (*UserDTO, error)
	GetUsers(ctx context.Context, itemsPerPage uint, currentPage uint) (PaginationDTO, error)
	UpdateUser(ctx context.Context, id uint, user schemas.User) (schemas.User, error)
}

type TokenPasswordRepository interface {
	CreateToken(ctx context.Context, token schemas.TokenPassword)
	UpdateToken(ctx context.Context, id uint, token schemas.TokenPassword)
	GetToken(ctx context.Context, id uint)
}

type ProductRepository interface {
	CreateProduct(ctx context.Context, product schemas.Product)
	GetProducts(ctx context.Context, itemsPerPage uint, currentPage uint)
	UpdateProducts(ctx context.Context, id uint)
}
