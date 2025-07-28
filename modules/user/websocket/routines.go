package websocket

import (
	"automatic-doodle/types"
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/websocket/v2"
)

func (c *Client) unregisterSafe(hub *Hub) {
	c.once.Do(func() {
		hub.unregister <- c
	})
}
func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			hub.clients[client.user.Id] = client
		case client := <-hub.unregister:
			delete(hub.clients, client.user.Id)
			close(client.send)
			client.conn.Close()
		case msg := <-hub.privateMessage:
			err := hub.websocketService.CreateMessage(msg, context.Background())
			if err != nil {
				if sender, ok := hub.clients[msg.From.Id]; ok {
					fmt.Println("Error in sending message:", err)
					sender.send <- msg
				}
			}

			if receiver, ok := hub.clients[msg.To]; ok {
				fmt.Println("Message sent to client", msg)
				receiver.send <- msg
			}
		}

	}
}

func (c *Client) Read(hub *Hub) {
	defer func() {
		c.unregisterSafe(hub)

	}()

	for {
		messageType, msg, err := c.conn.ReadMessage()

		if err != nil {
			fmt.Println(err)
			break
		}

		if messageType != websocket.TextMessage {
			fmt.Println("Unsupoorted Message Wont be send")
			errMsg := websocket.FormatCloseMessage(websocket.CloseUnsupportedData, "unsupported type")
			c.conn.WriteMessage(websocket.CloseMessage, errMsg)
			break
		}

		var pm types.PrivateMessage
		if err := json.Unmarshal(msg, &pm); err != nil {
			fmt.Println("Error unmarshalling message:", err)
			errMsg := websocket.FormatCloseMessage(websocket.CloseUnsupportedData, "unsupported type")
			c.conn.WriteMessage(websocket.CloseMessage, errMsg)
			break
		}

		pm.From = c.user

		fmt.Println("Private message")
		fmt.Println(pm.From)
		fmt.Println(pm.Message)
		fmt.Println(pm.To)
		hub.privateMessage <- &pm

	}

}

func (c *Client) Write(hub *Hub) {
	defer c.unregisterSafe(hub)

	for {
		select {
		case msg, ok := <-c.send:
			if !ok {
				return
			}

			fmt.Println("Sending message:", msg)

			if err := c.conn.WriteJSON(msg); err != nil {
				fmt.Println("Write error:", err)
				return
			}

		case msg := <-c.err:
			c.conn.WriteJSON("Error in sending message" + msg)
			return
		}
	}
}
