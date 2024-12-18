package factory

import (
	"automatic-doodle/ent"
	"sync"
)

type Factory struct {
	db ent.Client
}

var (
	module     Factory
	moduleOnce sync.Once
)

func New(db ent.Client) *Factory {
	moduleOnce.Do(func() {
		module = Factory{
			db: db}
	})
	return &module
}
