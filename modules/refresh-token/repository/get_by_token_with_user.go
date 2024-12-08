package repository

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/refreshtoken"
	"context"
	"fmt"
)

func (repo *Repository) GetByTokenWithUser(
	token string,
	ctx context.Context,
) (*ent.RefreshToken, error) {
	if len(token) == 0 {
		return nil, fmt.Errorf("token cannot be empty")
	}

	itemRow, err := repo.db.RefreshToken.
		Query().
		Where(refreshtoken.Token(token)).
		WithUsers().
		First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("failed to fetch refreshToken: %w", err)
	}

	return itemRow, nil
}
