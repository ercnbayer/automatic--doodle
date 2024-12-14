package di

import (
	fileFactory "automatic-doodle/modules/file/factory"
	fileRepository "automatic-doodle/modules/file/repository"
	fileRest "automatic-doodle/modules/file/rest"
	fileService "automatic-doodle/modules/file/service"
	router "automatic-doodle/modules/router"
	userService "automatic-doodle/modules/user/service"

	"automatic-doodle/pkg/logger"

	"github.com/google/wire"
)

var (
	FileFactoryProvider wire.ProviderSet = wire.NewSet(

		fileFactory.New,
		wire.Bind(new(fileService.FileFactory), new(*fileFactory.Factory)),
		wire.Bind(new(userService.FileFactory), new(*fileFactory.Factory)),
	)

	FileRepositoryProvider wire.ProviderSet = wire.NewSet(fileRepository.New,
		wire.Bind(new(fileService.FileRepository), new(*fileRepository.Repository)),
		wire.Bind(new(userService.FileRepository), new(*fileRepository.Repository)))
	FileRestProviderSet wire.ProviderSet = wire.NewSet(
		fileRest.New,

		wire.Bind(
			new(router.FileHandler),
			new(*fileRest.Rest),
		),

		wire.InterfaceValue(
			new(fileRest.Logger),
			logger.New("FileRest"),
		),
	)
	FileServiceProviderSet wire.ProviderSet = wire.NewSet(
		fileService.New,

		wire.Bind(
			new(fileRest.FileService),
			new(*fileService.Service),
		),

		wire.InterfaceValue(
			new(fileService.Logger),
			logger.New("FileService"),
		),
	)
)
