package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	"github.com/jeffreycharters/blurt/internal/blurthub"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

func (h *DefaultHandler) WebsocketHandler(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer func() {
			blurthub.Unregister <- ws
			ws.Close()
		}()

		// Register the client
		blurthub.Register <- ws

		for {
			var message blurthub.Message
			err := websocket.JSON.Receive(ws, &message)
			if err != nil && errors.Is(err, io.EOF) {
				slog.Info("connection closed")
				return
			}

			if err != nil {
				slog.Info("error reading message", "read error:", err)
				return
			}

			if message.ContentType == "" {
				slog.Info("no content type")
				return
			}

			var outgoing []byte
			if message.ContentType == blurthub.Blurt {
				var dest blurthub.BlurtRequired
				if err := json.Unmarshal(message.Payload, &dest); err != nil {
					websocket.JSON.Send(ws, err)
					return
				}

				user, err := h.db.SetUser(dest.Username)
				if err != nil {
					slog.Info("error setting user", "error", err)
					websocket.JSON.Send(ws, err)
					return
				}

				added_blurt, err := h.db.SetBlurt(user.Username, dest.Content)
				if err != nil {
					slog.Error("error setting blurt", "error", err)
					websocket.JSON.Send(ws, err)
					return
				}

				outgoing, err = json.Marshal(added_blurt)
				if err != nil {
					slog.Info("error marshalling outgoing message", "error", err)
					websocket.JSON.Send(ws, err)
					return
				}
			}

			if message.ContentType == blurthub.Lik {
				var dest blurthub.LikRequired
				if err := json.Unmarshal(message.Payload, &dest); err != nil {
					slog.Info("error unmarshalling incoming message", "error", err)
					websocket.JSON.Send(ws, err)
					return
				}

				if err := h.db.LikBlurt(dest.Likker, dest.BlurtID); err != nil {
					slog.Info("error liking blurt", "error", err)
					websocket.JSON.Send(ws, err)
					return
				}

				slog.Info("broadbasting")
				outgoing = message.Payload
			}

			blurthub.Broadcast <- blurthub.Message{
				ContentType: message.ContentType,
				Payload:     outgoing,
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil

}
