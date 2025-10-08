package schemas

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
