package service

import (
	"automatic-doodle/types"
	"sync"
)

var (
	module     Service
	moduleOnce sync.Once
)

type Service struct {
	cfg           ConfigModule
	log           Logger
	adminCreds    types.JWTCredentials
	customerCreds types.JWTCredentials
}

func New(
	cfg ConfigModule,
	log Logger,
) *Service {
	moduleOnce.Do(func() {
		module = Service{
			cfg: cfg,
			log: log,
		}

		module.Initialize()
	})

	return &module
}
