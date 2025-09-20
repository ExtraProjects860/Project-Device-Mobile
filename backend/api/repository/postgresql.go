package repository

type UserPostgreSQL interface {
	CreateUser()
	GetInfoUser()
	GetUsers()
	UpdateUser()
}

type RolePostgreSQL interface {
	CreateRole()
	GetRoles()
	UpdateRole()
}

type EnterprisePostgreSQL interface {
	CreateEnterprise()
	GetEnterprises()
	UpdateEnterprise()
}

type ProductPostgreSQL interface {
	CreateProduct()
	GetProducts()
	UpdateProducts()
}

type TokenPassword interface {
	CreateToken()
	UpdateToken()
	GetToken()
}

type WishList interface {
	AddInWishList()
	UpdateWishList()
	GetItensWishList()
}
