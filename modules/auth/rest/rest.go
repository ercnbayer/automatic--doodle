package rest

import "sync"

type Rest struct {
	log                      Logger
	authenticationService    AuthenticationService
	authenticationMiddleware AuthenticationMiddleware
}

var (
	module     Rest
	moduleOnce sync.Once
)

func New(
	log Logger,
	authenticationService AuthenticationService,
	authenticationMiddleware AuthenticationMiddleware,
) *Rest {
	moduleOnce.Do(func() {
		module = Rest{
			log:                      log,
			authenticationService:    authenticationService,
			authenticationMiddleware: authenticationMiddleware,
		}
	})

	return &module
}
