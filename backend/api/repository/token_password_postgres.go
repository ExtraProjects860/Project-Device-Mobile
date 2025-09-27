package repository

import (
	"context"
	"time"

	"github.com/ExtraProjects860/Project-Device-Mobile/schemas"
)

type TokenPasswordDTO struct {
	ID     uint       `json:"id"`
	Code   *string    `json:"code,omitempty"`
	TimeUp *time.Time `json:"time_up,omitempty"`
}

func makeTokenPasswordOutput(token schemas.TokenPassword) *TokenPasswordDTO {
	return &TokenPasswordDTO{
		ID:     token.ID,
		Code:   token.Code,
		TimeUp: token.TimeUp,
	}
}

func (r *postgresTokenPasswordRepository) CreateToken(ctx context.Context, token schemas.TokenPassword) {
	return
}

func (r *postgresTokenPasswordRepository) UpdateToken(ctx context.Context, id uint, token schemas.TokenPassword) {
	return
}

func (r *postgresTokenPasswordRepository) GetToken(ctx context.Context, id uint) {
	return
}
