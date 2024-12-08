package di

import (
	authService "automatic-doodle/modules/auth/service"
	"automatic-doodle/modules/router"
	UserFactory "automatic-doodle/modules/user/factory"
	UserRepo "automatic-doodle/modules/user/repository"
	UserRest "automatic-doodle/modules/user/rest"

	"github.com/google/wire"
)

var (
	UserFactoryProviderSet wire.ProviderSet = wire.NewSet(
		UserFactory.New,
		wire.Bind(new(authService.UserFactory), new(*UserFactory.UserFactory)),
	)
	UserRepositoryProviderSet wire.ProviderSet = wire.NewSet(UserRepo.New,
		wire.Bind(new(authService.UserRepository), new(*UserRepo.UserRepository)))

	UserRestProviderSet wire.ProviderSet = wire.NewSet(UserRest.New,

		wire.Bind(new(router.UserHandler), new(*UserRest.Rest)))
)
