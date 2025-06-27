package websocket

import (
	"automatic-doodle/types"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func (w *Websocket) WebSocketHandler(hub *Hub) fiber.Handler {
	return websocket.New(func(conn *websocket.Conn) {
		val := conn.Locals("pUser")

		if val == nil {
			conn.Close()
		}

		authUser, ok := val.(types.AuthenticatedUser)
		if !ok {
			log.Print("NOT A VALID USER")
			conn.Close()
			return
		}

		client := &Client{
			conn: conn,
			user: authUser,
			send: make(chan []byte, 4096),
		}
		hub.register <- client

		go client.Write(hub)
		client.Read(hub)
	})
}
