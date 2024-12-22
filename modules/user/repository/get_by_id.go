package repository

import (
	"automatic-doodle/ent"
	"context"

	"github.com/google/uuid"
)

func (repo *UserRepository) GetById(id uuid.UUID, ctx context.Context) (*ent.User, error) {

	user, err := repo.db.User.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, err
	}

	return user, nil
}
