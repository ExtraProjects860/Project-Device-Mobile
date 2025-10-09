package request

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type WishListRequest struct {
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
}

func (w *WishListRequest) Validate(validate *validator.Validate) error {
	if err := validate.Var(w.UserID, "required,gt=0"); err != nil {
		return fmt.Errorf("user_id: %v", err)
	}

	if err := validate.Var(w.ProductID, "required,gt=0"); err != nil {
		return fmt.Errorf("product_id: %v", err)
	}

	return nil
}

func (w *WishListRequest) ValidateUpdate(validate *validator.Validate) error {
	if w.UserID != 0 {
		if err := validate.Var(w.UserID, "gt=0"); err != nil {
			return fmt.Errorf("user_id: %v", err)
		}
	}

	if w.ProductID != 0 {
		if err := validate.Var(w.ProductID, "gt=0"); err != nil {
			return fmt.Errorf("product_id: %v", err)
		}
	}

	return nil
}

func (w *WishListRequest) Format() {}
