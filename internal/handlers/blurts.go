package handlers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

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
