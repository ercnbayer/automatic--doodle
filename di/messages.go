package di

import (
	"automatic-doodle/modules/messages/factory"
	"automatic-doodle/modules/messages/repository"
	"automatic-doodle/modules/messages/rest"
	"automatic-doodle/modules/messages/service"
	"automatic-doodle/modules/router"
	websocketService "automatic-doodle/modules/user/websocket/service"

	"github.com/google/wire"
)

var MessageFactoryProviderSet wire.ProviderSet = wire.NewSet(factory.New,
	wire.Bind(
		new(websocketService.MessageFactory),
		new(*factory.Factory),
	),
)

var MessageRepositoryProviderSet wire.ProviderSet = wire.NewSet(repository.New,
	wire.Bind(
		new(websocketService.MessageRepository),
		new(*repository.Repository),
	),
	wire.Bind(
		new(service.MessageRepository),
		new(*repository.Repository),
	),
)

var MessageServiceProviderSet wire.ProviderSet = wire.NewSet(service.New,
	wire.Bind(
		new(rest.MessageService),
		new(*service.Service),
	),
)
var MessageRestProviderSet wire.ProviderSet = wire.NewSet(rest.New,
	wire.Bind(
		new(router.MessageHandler),
		new(*rest.Rest),
	),
)
