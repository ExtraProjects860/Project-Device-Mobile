package dto

import (
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
