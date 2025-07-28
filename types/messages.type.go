package types

import (
	"automatic-doodle/ent"
	"time"

	"github.com/google/uuid"
)

type PrivateMessage struct {
	From    AuthenticatedUser `json:"from"`
	To      uuid.UUID         `json:"to"`
	Message string            `json:"message"`
}

type ReturnMessage struct {
	From      uuid.UUID `json:"from"`
	To        uuid.UUID `json:"to"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func EntToReturnMessage(message *ent.Messages) *ReturnMessage {

	return &ReturnMessage{
		From:      message.ID,
		To:        message.To,
		Message:   message.Message,
		Timestamp: message.CreatedAt,
	}
}
