package types

import (
	"automatic-doodle/ent/user"
	"time"

	"github.com/google/uuid"
)

type JWTTokenPayload struct {
	UserId   uuid.UUID `json:"userId"`
	UserRole user.Role `json:"userRole"`
}

type JWTCredentials struct {
	Key            string
	Aud            string
	ExpireDuration time.Duration

	UserId uuid.UUID
}
