package websocket

import (
	"automatic-doodle/ent"
	"automatic-doodle/types"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetById(
		uuid.UUID,
		context.Context,
	) (*ent.User, error)
}
type AuthenticationMiddleware interface {
	WebSocketAuth(c *fiber.Ctx) error
}

type WebsocketService interface {
	CreateMessage(
		*types.PrivateMessage,
		context.Context,
	) error
}
