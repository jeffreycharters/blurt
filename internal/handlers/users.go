package handlers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *DefaultHandler) NewUserHandler(c echo.Context) error {
	username := c.Param("username")

	user, err := h.db.SetUser(username)
	if err != nil {
		slog.Info("error setting user", "error", err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (h *DefaultHandler) UserCountHandler(c echo.Context) error {
	count, err := h.db.GetUserCount()
	if err != nil {
		slog.Info("error getting user count", "error", err)
		return err
	}

	return c.JSON(http.StatusOK, count)
}
