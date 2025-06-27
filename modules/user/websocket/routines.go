package websocket

import (
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
			if receiver, ok := hub.clients[msg.To]; ok {
				fmt.Println("Message sent to client", msg)
				jsonMsg, _ := json.Marshal(msg)
				receiver.send <- jsonMsg
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

		var pm PrivateMessage
		if err := json.Unmarshal(msg, &pm); err != nil {
			fmt.Println("Error unmarshalling message:", err)
			errMsg := websocket.FormatCloseMessage(websocket.CloseUnsupportedData, "unsupported type")
			c.conn.WriteMessage(websocket.CloseMessage, errMsg)
			break
		}

		pm.From = c.user

		hub.privateMessage <- &pm

	}

}

func (c *Client) Write(hub *Hub) {

	defer c.unregisterSafe(hub)

	for msg := range c.send {

		fmt.Println("Sending message:", string(msg))
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}
