package rest

import "sync"

type Rest struct {
	authenticationMiddleware AuthenticationMiddleware
	logger                   Logger
	jobFactory               JobFactory
	jobRepository            JobRepository
}

var (
	module     Rest
	moduleOnce sync.Once
)

func New(
	authenticationMiddleware AuthenticationMiddleware,
	jobFactory JobFactory,
	jobRepository JobRepository,
	logger Logger,
) *Rest {
	moduleOnce.Do(func() {
		module = Rest{
			authenticationMiddleware: authenticationMiddleware,
			logger:                   logger,
			jobFactory:               jobFactory,
			jobRepository:            jobRepository,
		}
	})

	return &module
}
