package handlers

import (
	"net/http"

	"github.com/jeffreycharters/blurt/internal/database"
	"github.com/labstack/echo/v4"
)

type DefaultHandler struct {
	db     *database.HandlerDB
	Routes *echo.Echo
}

func New(db *database.HandlerDB) DefaultHandler {
	return DefaultHandler{
		db: db,
	}
}

func (h *DefaultHandler) NewUserHandler(c echo.Context) error {
	username := c.Param("username")

	user, err := h.db.SetUser(username)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (h *DefaultHandler) UserCountHandler(c echo.Context) error {
	count, err := h.db.GetUserCount()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, count)
}
