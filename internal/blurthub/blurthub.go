package blurthub

import (
	"encoding/json"
	"log"
	"log/slog"
	"sync"

	"golang.org/x/net/websocket"
)

type Client struct {
	isClosing bool
	mu        sync.Mutex
}

type ContentType string

const (
	Blurt ContentType = "blurt"
	Lik   ContentType = "lik"
	Error ContentType = "error"
)

type BlurtRequired struct {
	Username string `json:"username" validate:"min=1,max=14,required"`
	Content  string `json:"content" validate:"min=1,max=14,required"`
}

type LikRequired struct {
	Likker  string `json:"likker" validate:"min=1,max=14,required"`
	BlurtID string `json:"blurt_id" validate:"uuid4,required"`
}

type Message struct {
	ContentType ContentType     `json:"content_type"`
	Payload     json.RawMessage `json:"payload"`
}

var Clients = make(map[*websocket.Conn]*Client)
var Register = make(chan *websocket.Conn)
var Broadcast = make(chan Message)
var Unregister = make(chan *websocket.Conn)

func Run() {
	for {
		select {
		case connection := <-Register:
			Clients[connection] = &Client{}
			slog.Info("connection registered", "count", len(Clients))

		case message := <-Broadcast:
			BroadcastAll(Clients, message)

		case connection := <-Unregister:
			// Remove the client from the pool
			delete(Clients, connection)
			slog.Info("connection unregistered", "count", len(Clients))
		}
	}
}

func BroadcastAll(clients map[*websocket.Conn]*Client, message Message) error {
	slog.Info("message broadcast", "count", len(clients), "type", message.ContentType, "payload", string(message.Payload))
	// Send the message to all clients
	for connection, c := range clients {
		go func(connection *websocket.Conn, c *Client) { // send to each client in parallel so we don't block on a slow client
			c.mu.Lock()
			defer c.mu.Unlock()
			if c.isClosing {
				return
			}

			if err := websocket.JSON.Send(connection, message); err != nil {
				c.isClosing = true
				log.Println("write error:", err)

				connection.Close()
				Unregister <- connection
			}
		}(connection, c)
	}

	return nil
}
