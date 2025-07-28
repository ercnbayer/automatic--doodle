package websocket

import (
	"automatic-doodle/types"
	"sync"

	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

type Client struct {
	conn *websocket.Conn
	user types.AuthenticatedUser
	send chan *types.PrivateMessage
	once sync.Once
	err  chan string
}

type Hub struct {
	register         chan *Client
	unregister       chan *Client
	privateMessage   chan *types.PrivateMessage
	clients          map[uuid.UUID]*Client
	websocketService WebsocketService
}

func NewHub(websocketService WebsocketService) *Hub {
	return &Hub{
		clients:          make(map[uuid.UUID]*Client),
		register:         make(chan *Client),
		unregister:       make(chan *Client),
		privateMessage:   make(chan *types.PrivateMessage),
		websocketService: websocketService,
	}
}
