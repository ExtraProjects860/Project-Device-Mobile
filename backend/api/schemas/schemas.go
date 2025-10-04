package schemas

import (
	"context"

	"github.com/go-playground/validator/v10"
)

type ValidateModel interface {
	Format()
	Validate(ctx context.Context, v *validator.Validate) error
}

func AllModelsSlice() []any {
	return []any{
		&User{},
		&TokenPassword{},
		&Enterprise{},
		&Role{},
		&Product{},
		&WishList{},
	}
}
