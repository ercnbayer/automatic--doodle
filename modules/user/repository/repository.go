package repository

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/user"
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

var (
	repository     UserRepository
	repositoryOnce sync.Once
)

type UserRepository struct {
	db *ent.Client
}

var (
	module     UserRepository
	moduleOnce sync.Once
)

func New(
	db *ent.Client,
) *UserRepository {
	moduleOnce.Do(func() {
		module = UserRepository{
			db: db,
		}
	})

	return &module
}

func (repo *UserRepository) Create(
	item *ent.UserCreate,
	ctx context.Context,
) (*ent.User, error) {

	if item == nil {
		return nil, fmt.Errorf("user cannot be nil")
	}

	itemRow, err := item.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return itemRow, nil
}

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
