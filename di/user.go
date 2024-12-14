package di

import (
	authService "automatic-doodle/modules/auth/service"
	pprest "automatic-doodle/modules/profile_photo/rest"
	"automatic-doodle/modules/router"
	UserFactory "automatic-doodle/modules/user/factory"
	UserRepo "automatic-doodle/modules/user/repository"
	UserRest "automatic-doodle/modules/user/rest"
	UserService "automatic-doodle/modules/user/service"
	"automatic-doodle/pkg/logger"

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

	UserServiceProviderSet wire.ProviderSet = wire.NewSet(UserService.New, wire.Bind(new(pprest.UserService), new(*UserService.UserService)),
		wire.Bind(new(pprest.UserService), new(*UserService.UserService)), wire.InterfaceValue(new(UserService.Logger), logger.New("UserService")),
	)
)
