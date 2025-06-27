package di

import (
	"automatic-doodle/modules/router"
	"automatic-doodle/modules/user/websocket"

	"github.com/google/wire"
)

var WebsocketProviderSet wire.ProviderSet = wire.NewSet(websocket.New,

	wire.Bind(new(router.WebsocketHandler), new(*websocket.Websocket)))
