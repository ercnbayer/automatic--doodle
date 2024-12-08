package repository

import (
	"context"
	"fmt"

	"automatic-doodle/ent"
	"automatic-doodle/ent/refreshtoken"
)

func (repo *Repository) Expire(
	token string,
	ctx context.Context,
) error {
	if len(token) == 0 {
		return fmt.Errorf("token cannot be empty")
	}

	err := repo.db.RefreshToken.
		Update().
		Where(refreshtoken.Token(token)).
		SetIsClaimed(true).
		Exec(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return fmt.Errorf("failed to fetch refreshToken: %w", err)
	}

	return nil
}
