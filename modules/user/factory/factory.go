package factory

import (
	"automatic-doodle/ent"
	"sync"
)

type UserFactory struct {
	db         *ent.Client
	encryption Encryption
}

var (
	module     UserFactory
	moduleOnce sync.Once
)

func New(
	db *ent.Client,
	encryption Encryption,
) *UserFactory {
	moduleOnce.Do(func() {
		module = UserFactory{
			db:         db,
			encryption: encryption,
		}
	})

	return &module
}
