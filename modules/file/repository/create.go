package repository

import (
	"automatic-doodle/ent"
	"context"
	"fmt"
)

func (r *Repository) Create(
	item *ent.FileCreate,
	ctx context.Context,
) (*ent.File, error) {
	if item == nil {
		return nil, fmt.Errorf("user cannot be nil")
	}

	itemRow, err := item.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return itemRow, nil
}
