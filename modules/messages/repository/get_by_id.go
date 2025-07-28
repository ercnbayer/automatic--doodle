package repository

import (
	"automatic-doodle/ent"
	"context"

	"github.com/google/uuid"
)

func (r *Repository) GetById(id uuid.UUID, ctx context.Context) (*ent.Messages, error) {

	return r.db.Messages.Get(ctx, id)
}
