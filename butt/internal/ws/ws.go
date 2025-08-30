package ws

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	isClosing bool
	mu        sync.Mutex
}

type Payload struct {
	Success bool        `json:"success"`
	Payload interface{} `json:"payload"`
}

var clients = make(map[*websocket.Conn]*Client)
var Register = make(chan *websocket.Conn)
var Broadcast = make(chan []byte)
var Unregister = make(chan *websocket.Conn)

func BlurtHub() {
	for {
		select {
		case connection := <-Register:
			clients[connection] = &Client{}
			updateUserCount(len(clients))
			log.Println("connection registered")

		case message := <-Broadcast:
			BroadcastAll(clients, message)

		case connection := <-Unregister:
			// Remove the client from the hub
			delete(clients, connection)
			updateUserCount(len(clients))

			log.Println("connection unregistered", len(clients))
		}
	}
}

func BroadcastAll(clients map[*websocket.Conn]*Client, message []byte) error {
	log.Println("message broadcasting:", string(message))
	// Send the message to all clients
	for connection, c := range clients {
		go func(connection *websocket.Conn, c *Client) { // send to each client in parallel so we don't block on a slow client
			c.mu.Lock()
			defer c.mu.Unlock()
			if c.isClosing {
				return
			}
			if err := connection.WriteMessage(websocket.TextMessage, message); err != nil {
				c.isClosing = true
				log.Println("write error:", err)

				connection.WriteMessage(websocket.CloseMessage, []byte{})
				connection.Close()
				Unregister <- connection
			}
		}(connection, c)
	}

	return nil
}

func updateUserCount(userCount int) error {
	payload := map[string]interface{}{
		"userCount": userCount,
	}
	userPayload, err := json.Marshal(Payload{
		Success: true,
		Payload: payload,
	})

	if err != nil {
		log.Println("error marshalling user payload:", err)
		return err
	}

	return BroadcastAll(clients, userPayload)
}
