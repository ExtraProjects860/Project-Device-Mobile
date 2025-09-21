package repository

type ProductPostgres interface {
	CreateProduct()
	GetProducts()
	UpdateProducts()
}
