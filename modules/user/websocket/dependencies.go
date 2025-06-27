package websocket

import (
	"automatic-doodle/ent"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetById(uuid.UUID, context.Context) (*ent.User, error)
}
type AuthenticationMiddleware interface {
	WebSocketAuth(c *fiber.Ctx) error
}
