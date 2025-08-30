package handlers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *DefaultHandler) NewBlurtHandler(c echo.Context) error {
	type req_body struct {
		Username string `json:"username"`
		Blurt    string `json:"blurt"`
	}

	var body req_body
	if err := c.Bind(&body); err != nil {
		slog.Info("error binding", "error", err)
		return err
	}

	if body.Username == "" || body.Blurt == "" {
		return echo.ErrBadRequest
	}

	user, err := h.db.SetUser(body.Username)
	if err != nil {
		slog.Info("error setting user", "error", err)
		return err
	}

	blurt, err := h.db.SetBlurt(user.Username, body.Blurt)
	if err != nil {
		slog.Info("error setting blurt", "error", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, blurt)
}

func (h *DefaultHandler) GetBlurtsHandler(c echo.Context) error {
	type query_params struct {
		Offset int `query:"offset"`
	}

	var params query_params
	_ = c.Bind(&params)

	blurts, err := h.db.GetBlurts(params.Offset)
	if err != nil {
		slog.Info("error getting blurts", "error", err)
		return err
	}

	return c.JSON(http.StatusOK, blurts)

}
