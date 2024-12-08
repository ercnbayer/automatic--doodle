package repository

import (
	"sync"

	"automatic-doodle/ent"
)

var (
	module     Repository
	moduleOnce sync.Once
)

func New(
	db *ent.Client,
) *Repository {
	moduleOnce.Do(func() {
		module = Repository{
			db: db,
		}
	})

	return &module
}

type Repository struct {
	db *ent.Client
}
