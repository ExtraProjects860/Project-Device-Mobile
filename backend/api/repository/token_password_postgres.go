package repository

type TokenPasswordPostgres interface {
	CreateToken()
	UpdateToken()
	GetToken()
}
