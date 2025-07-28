package repository

import (
	"automatic-doodle/ent"
	"context"
	"fmt"
)

func (r *Repository) Create(item *ent.MessagesCreate, ctx context.Context) (*ent.Messages, error) {
	if item == nil {
		return nil, fmt.Errorf("Message cannot be nil")
	}

	itemRow, err := item.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create Message: %w", err)
	}

	return itemRow, nil
}
