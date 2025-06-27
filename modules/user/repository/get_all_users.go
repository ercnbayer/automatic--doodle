package repository

import (
	"automatic-doodle/ent"
	"context"
)

func (repo *UserRepository) GetAllUsers(ctx context.Context) ([]*ent.User, error) {
	users, err := repo.db.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
} 