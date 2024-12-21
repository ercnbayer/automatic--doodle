package rest

import "sync"

type Rest struct {
	jobApplService JobApplService
	logger         Logger
	authMiddleware AuthMiddleware
}

var (
	module     Rest
	moduleOnce sync.Once
)

func New(jobApplService JobApplService, logger Logger, authMiddleware AuthMiddleware) *Rest {
	moduleOnce.Do(func() {
		module = Rest{
			jobApplService: jobApplService,
			logger:         logger,
			authMiddleware: authMiddleware,
		}
	})
	return &module
}
