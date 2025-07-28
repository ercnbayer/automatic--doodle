package service

import (
	"automatic-doodle/ent"
	"context"

	"github.com/google/uuid"
)

type MessageRepository interface {
	GetMessagesByUserId(
		context.Context,
		uuid.UUID,
	) ([]*ent.Messages, error)

	FetchConversationsByUserId(
		uuid.UUID,
		uuid.UUID,
		context.Context,
	) ([]*ent.Messages, error)
}
