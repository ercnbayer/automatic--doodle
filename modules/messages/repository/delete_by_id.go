package repository

import (
	"context"

	"github.com/google/uuid"
)

func (r *Repository) DeleteById(id uuid.UUID, ctx context.Context) error {
	return r.db.Messages.DeleteOneID(id).Exec(ctx)
}
