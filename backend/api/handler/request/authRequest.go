package request

// JWT Token

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserData struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}
