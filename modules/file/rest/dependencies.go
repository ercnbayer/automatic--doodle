package rest

import (
	"automatic-doodle/types"

	"github.com/gofiber/fiber/v2"
)

type Logger interface {
	Trace(format string, args ...any)
	Info(format string, args ...any)
	Warning(format string, args ...any)
	Error(format string, args ...any)
	Fatal(format string, args ...any)
}

type FileService interface {
	CreateUploadUrl(
		*types.CreateUploadUrlRequest,
	) (types.CreateUploadUrlResponse, error)
}

type AuthenticationMiddleware interface {
	Authentication(*fiber.Ctx) error
}
