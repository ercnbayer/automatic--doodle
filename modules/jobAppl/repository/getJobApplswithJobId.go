package repository

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/jobapplication"
	"context"

	"github.com/google/uuid"
)

func (r *Repository) GetJobApplsWithJobId(offset int, limit int, jobID uuid.UUID, ctx context.Context) ([]*ent.JobApplication, error) {
	items, err := r.db.JobApplication.Query().Where(
		jobapplication.Or(
			jobapplication.JobID(jobID),
		)).WithJob().Offset(offset).Limit(limit).All(ctx)

	if err != nil {
		return nil, err
	}

	return items, nil
}
