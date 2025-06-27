package websocket

import (
	"automatic-doodle/types"
	"sync"

	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

type PrivateMessage struct {
	From    types.AuthenticatedUser `json:"from"`
	To      uuid.UUID               `json:"to"`
	Message string                  `json:"message"`
}

type Client struct {
	conn *websocket.Conn
	user types.AuthenticatedUser
	send chan []byte
	once sync.Once
}

type Hub struct {
	register       chan *Client
	unregister     chan *Client
	privateMessage chan *PrivateMessage
	clients        map[uuid.UUID]*Client
}

func NewHub() *Hub {
	return &Hub{
		clients:        make(map[uuid.UUID]*Client),
		register:       make(chan *Client),
		unregister:     make(chan *Client),
		privateMessage: make(chan *PrivateMessage),
	}
}
