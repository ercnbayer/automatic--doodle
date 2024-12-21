package rest

import (
	"automatic-doodle/types"

	"github.com/gofiber/fiber/v2"
)

type JobApplService interface {
	CreateJobAppl(request types.CreateJobApplRequest) (types.CreateJobApplResponse, error)
}

type Logger interface {
	Trace(format string, args ...any)
	Info(format string, args ...any)
	Warning(format string, args ...any)
	Error(format string, args ...any)
	Fatal(format string, args ...any)
}

type AuthMiddleware interface {
	Auth(*fiber.Ctx) error
}
