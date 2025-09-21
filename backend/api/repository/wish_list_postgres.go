package repository

type WishListPostgres interface {
	AddInWishList()
	UpdateWishList()
	GetItensWishList()
}
