package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	"github.com/jeffreycharters/blurt/internal/blurthub"
	"github.com/jeffreycharters/blurt/internal/validator"
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
				continue
			}

			if message.ContentType == "" {
				slog.Info("no content type")
				continue
			}

			validate := validator.GetValidator()

			var outgoing []byte
			if message.ContentType == blurthub.Blurt {
				var dest blurthub.BlurtRequired
				if err := json.Unmarshal(message.Payload, &dest); err != nil {
					sendError(ws, "bad request")
					continue
				}

				if err := validate.Struct(dest); err != nil {
					slog.Info("error validating incoming message", "error", err)
					sendError(ws, "username or blurt failed validation")
					continue
				}

				user, err := h.db.SetUser(dest.Username)
				if err != nil {
					slog.Info("error setting user", "error", err)
					sendError(ws, "internal error")
					continue
				}

				added_blurt, err := h.db.SetBlurt(user.Username, dest.Content)
				if err != nil {
					slog.Error("error setting blurt", "error", err)
					sendError(ws, "internal error")
					continue
				}

				outgoing, err = json.Marshal(added_blurt)
				if err != nil {
					slog.Info("error marshalling outgoing message", "error", err)
					sendError(ws, "database updated, reload page")
					continue
				}
			}

			if message.ContentType == blurthub.Lik {
				var dest blurthub.LikRequired
				if err := json.Unmarshal(message.Payload, &dest); err != nil {
					slog.Info("error unmarshalling incoming message", "error", err)
					sendError(ws, "bad request")
					continue
				}

				if err := validate.Struct(dest); err != nil {
					slog.Info("error validating incoming message", "error", err)
					sendError(ws, "username or blurt failed validation")
					continue
				}

				if err := h.db.LikBlurt(dest.Likker, dest.BlurtID); err != nil {
					slog.Info("error liking blurt", "error", err)
					sendError(ws, "error likking blurt")
					continue
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

func sendError(ws *websocket.Conn, message string) error {
	payload, err := json.Marshal(struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
	if err != nil {
		slog.Info("error marshalling error", "error", err)
		return err
	}

	return websocket.JSON.Send(ws, blurthub.Message{
		ContentType: blurthub.Error,
		Payload:     payload,
	})
}
