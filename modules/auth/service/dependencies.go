package service

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/user"
	"automatic-doodle/types"
	"context"

	"github.com/google/uuid"
)

type Encryption interface {
	EncryptTokens(
		tokens types.EncryptedTokenPayload,
	) (string, error)
	DecryptTokens(string) (types.EncryptedTokenPayload, error)
	CheckPasswordHash(string, string, string) bool
}
type AccessTokenService interface {
	Create(*types.JWTTokenPayload) (string, error)
	Validate(user.Role, *string) (string, error)
}

type Logger interface {
	Trace(format string, args ...any)
	Info(format string, args ...any)
	Warning(format string, args ...any)
	Error(format string, args ...any)
	Fatal(format string, args ...any)
}

type UserFactory interface {
	Create(
		phoneNumber,
		email,
		firstName,
		lastName,

		password string,
		role user.Role,
		state user.State,
	) *ent.UserCreate
}

type UserRepository interface {
	Create(*ent.UserCreate, context.Context) (*ent.User, error)
	GetById(uuid.UUID, context.Context) (*ent.User, error)
	DeleteUser(uuid.UUID, context.Context) error
	GetByIdentifier(string, context.Context) (*ent.User, error)
}

type AccessTokenRepository interface {
	Validate(user.Role, *string) (string, error)
}
type RefreshTokenFactory interface {
	Create(uuid.UUID) *ent.RefreshTokenCreate
}

type RefreshTokenRepository interface {
	Create(
		*ent.RefreshTokenCreate,
		context.Context,
	) (*ent.RefreshToken, error)
	GetByTokenWithUser(string, context.Context) (*ent.RefreshToken, error)
}
