package types

type LoginRequest struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"password,required"`
}

type TokenResponse struct {
	Token string            `json:"token"`
	User  AuthenticatedUser `json:"user"`
}
