package types

type LoginRequest struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"password,required"`
}

type TokenResponse struct {
	Token string            `json:"token"`
	User  AuthenticatedUser `json:"user"`
}

type RegisterRequest struct {
	FirstName   string `json:"firstName" validate:"required,min=3,max=255"`
	LastName    string `json:"lastName" validate:"required,min=3,max=255"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phoneNumber" validate:"required,e164"`
	Password    string `json:"password" validate:"required,password"`
}
