package types

import (
	"time"

	"github.com/google/uuid"
)

type JWTTokenPayload struct {
	UserId uuid.UUID `json:"userId"`
}

type JWTCredentials struct {
	Key            string
	Aud            string
	ExpireDuration time.Duration

	UserId uuid.UUID
}
