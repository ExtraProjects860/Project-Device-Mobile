package repository

type RolePostgres interface {
	CreateRole()
	GetRoles()
	UpdateRole()
}
