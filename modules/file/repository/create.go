package repository

import (
	"automatic-doodle/ent"
	"errors"

	"context"
)

func (r *Repository) Create(
	item *ent.FileCreate,
	ctx context.Context,
) (*ent.File, error) {
	if item == nil {
		return nil, errors.New("Null item")
	}

	itemRow, err := item.Save(ctx)
	if err != nil {
		return nil, err
	}

	return itemRow, nil
}
