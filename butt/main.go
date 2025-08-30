package main

import (
	"context"
	"flag"
	"log"
	"net/url"
	"strconv"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jeffreycharters/blurt/ent"
	"github.com/jeffreycharters/blurt/ent/blurt"
	"github.com/jeffreycharters/blurt/ent/user"
	"github.com/jeffreycharters/blurt/internal/db"
	"github.com/jeffreycharters/blurt/internal/ws"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	client := db.Client()
	defer client.Close()

	app.Post("/api/v1/users/:username", func(c *fiber.Ctx) error {
		username, err := url.QueryUnescape(c.Params("username"))
		if err != nil {
			return c.SendStatus(500)
		}

		user, err := client.User.Query().Where(user.Username(username)).First(context.Background())
		if ent.IsNotFound(err) {
			user, err = client.User.Create().SetUsername(username).Save(context.Background())
			if err != nil {
				return c.SendStatus(500)
			}
		} else if err != nil {
			return c.SendStatus(500)
		}

		return c.JSON(user)
	})

	app.Get("/api/v1/blurts", func(c *fiber.Ctx) error {
		offset, err := strconv.Atoi(c.Query("offset", "0"))
		if err != nil {
			offset = 0
		}

		count, err := strconv.Atoi(c.Query("count", "25"))
		if err != nil {
			count = 25
		}

		blurts, err := client.Blurt.Query().
			WithAuthor().
			WithLiks().
			Order(ent.Desc((blurt.FieldCreateTime))).
			Offset(offset).
			Limit(count).
			All(context.Background())
		if err != nil {
			return c.SendStatus(500)
		}

		return c.JSON(blurts)
	})

	app.Use(func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) { // Returns true if the client requested upgrade to the WebSocket protocol
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	go ws.BlurtHub()

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		// When the function returns, unregister the client and close the connection
		defer func() {
			ws.Unregister <- c
			c.Close()
		}()

		// Register the client
		ws.Register <- c

		for {
			messageType, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read error:", err)
				}

				return // Calls the deferred function, i.e. closes the connection on error
			}

			if messageType == websocket.TextMessage {
				// Broadcast the received message
				ret, err := db.ParseMessage(message)
				if err != nil {
					c.WriteJSON(db.WrapError(err))
					return
				}
				ws.Broadcast <- ret
			} else {
				log.Println("websocket message received of type", messageType)
			}
		}
	}))

	addr := flag.String("addr", ":8080", "http service address")
	flag.Parse()
	log.Fatal(app.Listen(*addr))
}
