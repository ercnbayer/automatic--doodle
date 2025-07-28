package repository

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/messages"
	"automatic-doodle/ent/user"
	"context"

	"github.com/google/uuid"
)

func (r *Repository) GetMessagesByUserId(
	ctx context.Context,
	userID uuid.UUID,
) ([]*ent.Messages, error) {

	filter := messages.Or(
		messages.HasSenderWith((user.ID(userID))),
		messages.HasReceiverWith(user.ID(userID)),
	)

	query, err := r.db.Messages.Query().
		Where(filter).
		Order(ent.Desc(messages.FieldCreatedAt)).WithReceiver().WithSender().All(ctx)

	if err != nil {
		return nil, err
	}

	return query, nil
}
