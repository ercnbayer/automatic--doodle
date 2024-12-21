package repository

import (
	"automatic-doodle/ent"
	"context"
	"fmt"
)

func (r *Repository) Create(item *ent.JobApplicationCreate, ctx context.Context) (*ent.JobApplication, error) {
	if item == nil {
		return nil, fmt.Errorf("jobAppl cannot be nil")
	}

	itemRow, err := item.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create JobAppl: %w", err)
	}

	return itemRow, nil
}
