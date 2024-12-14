package rest

import (
	"automatic-doodle/types"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserService interface {
	UpdateProfilePhoto(uuid.UUID, *types.UpdateProfilePhotoRequest) (types.AuthenticatedUser, error)
}

type AuthMiddleware interface {
	Auth(*fiber.Ctx) error
}

type Logger interface {
	Trace(format string, args ...any)
	Info(format string, args ...any)
	Warning(format string, args ...any)
	Error(format string, args ...any)
	Fatal(format string, args ...any)
}
