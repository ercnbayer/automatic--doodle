package websocket

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func (ws *Websocket) UpgradeToWebsocket(c *fiber.Ctx) error {
	fmt.Println("UpgradeToWebsocket")
	if websocket.IsWebSocketUpgrade(c) {
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}
