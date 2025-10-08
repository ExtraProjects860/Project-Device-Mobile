package schemas

/*
TODO
Remover validate e format de outros schemas e passar para request como struct seguindo modelo
do userRequest, além disso, passar a interface do validate para o request e anexar ela para uso
*/

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
