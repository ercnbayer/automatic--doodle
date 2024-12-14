package rest

import "sync"

var (
	module     Rest
	moduleOnce sync.Once
)

func New(
	log Logger,
	fileService FileService,
	authenticationMiddleware AuthenticationMiddleware,
) *Rest {
	moduleOnce.Do(func() {
		module = Rest{
			log:                      log,
			fileService:              fileService,
			authenticationMiddleware: authenticationMiddleware,
		}
	})

	return &module
}

type Rest struct {
	log                      Logger
	fileService              FileService
	authenticationMiddleware AuthenticationMiddleware
}
