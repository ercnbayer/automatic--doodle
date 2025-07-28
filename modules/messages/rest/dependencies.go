package rest

import (
	"automatic-doodle/types"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MessageService interface {
	GetMessagesByUserId(
		uuid.UUID,
		context.Context,
	) ([]types.ReturnMessage, error)
	FetchConversationsByUserId(
		uuid.UUID,
		uuid.UUID,
		context.Context,
	) ([]types.ReturnMessage, error)
}

type AuthenticationMiddleware interface {
	Auth(*fiber.Ctx) error
}
