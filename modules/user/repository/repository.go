package repository

import (
	"automatic-doodle/ent"
	"context"
	"fmt"
	"sync"
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
