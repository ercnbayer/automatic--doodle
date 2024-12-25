package rest

import (
	"automatic-doodle/ent"
	"automatic-doodle/types"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthenticationMiddleware interface {
	Auth(*fiber.Ctx) error
}

type Logger interface {
	Trace(format string, args ...any)
	Info(format string, args ...any)
	Warning(format string, args ...any)
	Error(format string, args ...any)
	Fatal(format string, args ...any)
}

type JobFactory interface {
	Create(int, string, string, time.Time, time.Time, uuid.UUID) *ent.JobCreate
}

type JobRepository interface {
	Create(*ent.JobCreate, context.Context) (*ent.Job, error)
}

type JobService interface {
	BrowseJob(offset int, limit int, identifier string, ctx context.Context) ([]types.JobAdvResponse, error)
}
