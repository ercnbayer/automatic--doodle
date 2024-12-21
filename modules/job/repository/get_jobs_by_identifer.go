package repository

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/job"
	"context"
)

func (repo *Repository) GetByIdentifier(
	identifier string,
	ctx context.Context,
) ([]*ent.Job, error) {
	items, err := repo.db.Job.Query().
		Where(
			job.Or(
				job.JobTypeContains(identifier),

				job.DescriptionContains(identifier),
			),
		).All(ctx)
	if err != nil {
		return nil, err
	}

	return items, nil
}
