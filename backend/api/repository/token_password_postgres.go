package repository

import "time"

type TokenPasswordPostgres interface {
	CreateToken()
	UpdateToken()
	GetToken()
}

type TokenPasswordDTO struct {
	ID     uint       `json:"id"`
	Code   *string    `json:"code,omitempty"`
	TimeUp *time.Time `json:"time_up,omitempty"`
}
