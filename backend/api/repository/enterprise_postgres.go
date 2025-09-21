package repository

type EnterprisePostgres interface {
	CreateEnterprise()
	GetEnterprises()
	UpdateEnterprise()
}
