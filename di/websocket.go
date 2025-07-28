package di

import (
	"automatic-doodle/modules/router"
	"automatic-doodle/modules/user/websocket"
	websocketService "automatic-doodle/modules/user/websocket/service"

	"github.com/google/wire"
)

var WebsocketProviderSet wire.ProviderSet = wire.NewSet(websocket.New,

	wire.Bind(new(router.WebsocketHandler), new(*websocket.Websocket)))

var WebsocketServiceProvider wire.ProviderSet = wire.NewSet(
	websocketService.New,
	wire.Bind(
		new(websocket.WebsocketService),
		new(*websocketService.WebsocketService),
	),
)
