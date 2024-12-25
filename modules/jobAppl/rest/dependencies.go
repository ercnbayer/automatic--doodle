package rest

import (
	"automatic-doodle/types"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type JobApplService interface {
	CreateJobAppl(request types.CreateJobApplRequest) (types.CreateJobApplResponse, error)
	CheckJobOwner(int, int, uuid.UUID, uuid.UUID, context.Context) ([]types.GetJobApplResponse, error)
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
