package websocket

import (
	"automatic-doodle/types"

	"github.com/gofiber/fiber/v2"
)

func (ws *Websocket) GetRoutes() []types.HandlerItem {
	return []types.HandlerItem{
		{
			Path:   "user/ws/",
			Method: "GET",
			Handler: []func(*fiber.Ctx) error{
				ws.authenticationMiddleware.WebSocketAuth,
				ws.UpgradeToWebsocket,
				ws.WebSocketHandler(ws.hub),
			},
		},
	}
}
