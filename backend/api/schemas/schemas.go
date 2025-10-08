package schemas

import (
	"context"

	"github.com/go-playground/validator/v10"
)

/*
TODO
Remover validate e format de outros schemas e passar para request como struct seguindo modelo
do userRequest, além disso, passar a interface do validate para o request e anexar ela para uso
*/

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
