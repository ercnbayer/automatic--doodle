package repository

import (
	"context"

	"github.com/google/uuid"
)

func (repo *UserRepository) DeleteUser(id uuid.UUID, ctx context.Context) error {
	err := repo.db.User.DeleteOneID(id).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
