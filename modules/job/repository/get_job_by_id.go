package repository

import (
	"automatic-doodle/ent"
	"context"

	"github.com/google/uuid"
)

func (r *Repository) GetJobById(id uuid.UUID, ctx context.Context) (*ent.Job, error) {

	return r.db.Job.Get(ctx, id)

}
