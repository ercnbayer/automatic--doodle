package repository

import (
	"automatic-doodle/ent"
	"sync"
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
