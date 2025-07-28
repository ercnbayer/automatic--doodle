package service

import (
	"automatic-doodle/ent"
	"context"

	"github.com/google/uuid"
)

type MessageFactory interface {
	Create(
		uuid.UUID,
		uuid.UUID,
		string,
	) *ent.MessagesCreate
}

type MessageRepository interface {
	Create(
		*ent.MessagesCreate,
		context.Context,
	) (*ent.Messages, error)
	GetById(
		uuid.UUID,
		context.Context,
	) (*ent.Messages, error)
}
