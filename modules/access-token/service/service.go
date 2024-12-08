package service

import (
	"automatic-doodle/pkg/config"
	"automatic-doodle/pkg/logger"
	"automatic-doodle/types"
	"sync"
)

var (
	module     Service
	moduleOnce sync.Once
)

type Service struct {
	cfg           config.ConfigModule
	log           logger.Logger
	adminCreds    types.JWTCredentials
	customerCreds types.JWTCredentials
}

func New(
	cfg config.ConfigModule,
	log logger.Logger,
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
