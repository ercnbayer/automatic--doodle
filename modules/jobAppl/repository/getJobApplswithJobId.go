package repository

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/jobapplication"
	"context"

	"github.com/google/uuid"
)

func (r *Repository) GetJobApplsWithJobId(jobID uuid.UUID, ctx context.Context) ([]*ent.JobApplication, error) {
	items, err := r.db.JobApplication.Query().Where(
		jobapplication.Or(
			jobapplication.JobID(jobID),
		)).All(ctx)

	if err != nil {
		return nil, err
	}

	return items, nil
}
