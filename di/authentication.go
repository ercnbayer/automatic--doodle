package di

import (
	authenticationMiddleware "automatic-doodle/modules/auth/middleware"
	authRest "automatic-doodle/modules/auth/rest"
	authService "automatic-doodle/modules/auth/service"
	FileRest "automatic-doodle/modules/file/rest"
	jobRest "automatic-doodle/modules/job/rest"
	ProfileRest "automatic-doodle/modules/profile/rest"
	"automatic-doodle/modules/router"
	userRest "automatic-doodle/modules/user/rest"

	logger "github.com/Wodemy-Labs/crawl"

	jobApplRest "automatic-doodle/modules/jobAppl/rest"

	"github.com/google/wire"
)

var (
	AuthenticationServiceProviderSet wire.ProviderSet = wire.NewSet(
		authService.New,
		wire.Bind(
			new(authenticationMiddleware.AuthenticationService),
			new(*authService.Service),
		),
		wire.Bind(new(authRest.AuthenticationService), new(*authService.Service)),

		wire.InterfaceValue(
			new(authService.Logger),
			logger.New("AuthenticationService", nil),
		))

	AuthenticationMiddlewareProviderSet wire.ProviderSet = wire.NewSet(
		authenticationMiddleware.New,

		wire.Bind(
			new(userRest.AuthenticationMiddleware),
			new(*authenticationMiddleware.Middleware),
		),
		wire.Bind(new(authRest.AuthenticationMiddleware), new(*authenticationMiddleware.Middleware)),
		wire.Bind(new(ProfileRest.AuthMiddleware), new(*authenticationMiddleware.Middleware)),
		wire.Bind(new(FileRest.AuthenticationMiddleware), new(*authenticationMiddleware.Middleware)),
		wire.Bind(new(jobRest.AuthenticationMiddleware), new(*authenticationMiddleware.Middleware)),
		wire.Bind(new(jobApplRest.AuthMiddleware), new(*authenticationMiddleware.Middleware)),

		wire.InterfaceValue(new(authenticationMiddleware.Logger), logger.New("Authentication Middleware Logger", nil)),
	)
	AuthenticationRestProviderSet wire.ProviderSet = wire.NewSet(authRest.New,
		wire.Bind(new(router.AuthHandler), new(*authRest.Rest)),

		wire.InterfaceValue(new(authRest.Logger), logger.New("AuthRestLogger", nil)))
)
