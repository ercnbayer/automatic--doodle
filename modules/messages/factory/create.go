package factory

import (
	"automatic-doodle/ent"

	"github.com/google/uuid"
)

func (f *Factory) Create(
	fromID uuid.UUID,
	toID uuid.UUID,
	message string,
) *ent.MessagesCreate {
	return f.db.Messages.Create().
		SetTo(toID).
		SetFrom(fromID).
		SetMessage(message)
}
