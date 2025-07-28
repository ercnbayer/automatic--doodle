package repository

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/messages"
	"automatic-doodle/ent/user"
	"context"

	"github.com/google/uuid"
)

func (r *Repository) FetchConversationsByUserId(
	userID uuid.UUID,
	withUserID uuid.UUID,
	ctx context.Context,
) ([]*ent.Messages, error) {

	filter := messages.Or(
		messages.And(
			messages.HasSenderWith((user.ID(userID))),
			messages.HasReceiverWith((user.ID(withUserID))),
		),
		messages.And(
			messages.HasReceiverWith(user.ID(userID)),
			messages.HasSenderWith(user.ID(withUserID)),
		),
	)

	query, err := r.db.Messages.Query().
		Where(filter).
		Order(ent.Desc(messages.FieldCreatedAt)).WithReceiver().WithSender().All(ctx)

	if err != nil {
		return nil, err
	}

	return query, nil
}
