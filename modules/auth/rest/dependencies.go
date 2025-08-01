package rest

import (
	"automatic-doodle/types"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Logger interface {
	Trace(format string, args ...any)
	Info(format string, args ...any)
	Warning(format string, args ...any)
	Error(format string, args ...any)
	Fatal(format string, args ...any)
}

type AuthenticationService interface {
	Login(*types.LoginRequest) (types.TokenResponse, error)
	Register(*types.RegisterRequest) (types.TokenResponse, error)
	UpdatePassword(req types.UpdateUserPassword, id uuid.UUID) (types.AuthenticatedUser, error)
}

type AuthenticationMiddleware interface {
	Auth(*fiber.Ctx) error
}
