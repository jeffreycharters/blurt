package handlers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *DefaultHandler) NewLikHandler(c echo.Context) error {
	type req_body struct {
		Username string `json:"username"`
		BlurtID  string `json:"blurt_id"`
	}

	var body req_body
	if err := c.Bind(&body); err != nil {
		slog.Info("error binding", "error", err)
		return err
	}

	if body.Username == "" || body.BlurtID == "" {
		return echo.ErrBadRequest
	}

	if _, err := h.db.SetUser(body.Username); err != nil {
		slog.Info("error setting user", "error", err)
		return err
	}

	err := h.db.LikBlurt(body.Username, body.BlurtID)
	if err != nil {
		slog.Info("error liking blurt", "error", err)
		return c.NoContent(http.StatusConflict)
	}

	return c.NoContent(http.StatusCreated)
}
