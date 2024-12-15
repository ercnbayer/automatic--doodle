package repository

import (
	"automatic-doodle/ent"
	"context"
	"fmt"
)

func (r *Repository) Create(item *ent.JobCreate, ctx context.Context) (*ent.Job, error) {

	if item == nil {
		return nil, fmt.Errorf("job pointer cannot be nil")
	}

	itemRow, err := item.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create job: %w", err)
	}

	return itemRow, nil
}
