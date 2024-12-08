package factory

import (
	"automatic-doodle/ent"
	"sync"
)

var (
	module     Factory
	moduleOnce sync.Once
)

func New(
	db *ent.Client,
	encryption Encryption,
) *Factory {
	moduleOnce.Do(func() {
		module = Factory{
			db:         db,
			encryption: encryption,
		}
	})

	return &module
}

type Factory struct {
	encryption Encryption
	db         *ent.Client
}
