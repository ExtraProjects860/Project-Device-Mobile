package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// JWT Token

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserData struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

func errParamIsRequired(name string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func getIdQuery(ctx *gin.Context) error {
	id := ctx.Query("id")
	if id == "" {
		return errParamIsRequired("id", "queryParameter")
	}

	return nil
}
