package repository

import (
	"automatic-doodle/ent"
	"sync"
)

type Repository struct {
	db ent.Client
}

var (
	module     Repository
	moduleOnce sync.Once
)

func New() *Repository {
	moduleOnce.Do(func() {
		module = Repository{}
	})
	return &module
}
