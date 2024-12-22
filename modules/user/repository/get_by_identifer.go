package repository

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/user"
	"context"
)

func (repo *UserRepository) GetByIdentifier(
	identifier string,
	ctx context.Context,
) (*ent.User, error) {
	item, err := repo.db.User.Query().
		Where(
			user.Or(
				user.Email(identifier),
				// user.Email(identifier),
				user.PhoneNumber(identifier),
			),
		).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return item, nil
}
