package types

import "automatic-doodle/ent/user"

type EncryptedTokenPayload struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	UserRole     user.Role `json:"userRole"`
}

type EncryptedTokenResult struct {
	AccessToken     string
	RefreshToken    string
	JwtTokenPayload JWTTokenPayload
}
