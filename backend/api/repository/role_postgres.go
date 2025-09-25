package repository

import "time"

type RolePostgres interface {
	CreateRole()
	GetRoles()
	UpdateRole()
}

type RoleDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
