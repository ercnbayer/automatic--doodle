package repository

import (
	"automatic-doodle/ent"
	"context"
	"fmt"
)

func (repo *Repository) Create(
	item *ent.RefreshTokenCreate,
	ctx context.Context,
) (*ent.RefreshToken, error) {
	if item == nil {
		return nil, fmt.Errorf("refreshToken cannot be nil")
	}

	itemRow, err := item.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create refreshToken: %w", err)
	}

	return itemRow, nil
}
