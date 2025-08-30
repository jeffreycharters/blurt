package handlers

import (
	"github.com/jeffreycharters/blurt/internal/database"
	"github.com/labstack/echo/v4"
)

type DefaultHandler struct {
	db     *database.HandlerDB
	router *echo.Echo
}

func New(db *database.HandlerDB) DefaultHandler {
	return DefaultHandler{
		db: db,
	}
}
