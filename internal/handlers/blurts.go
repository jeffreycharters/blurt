package handlers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

const BLURT_COUNT = 100

func (h *DefaultHandler) GetBlurtsHandler(c echo.Context) error {
	type query_params struct {
		Offset int64 `query:"offset"`
		Count  int64 `query:"count"`
	}

	var params query_params
	_ = c.Bind(&params)

	if params.Count == 0 {
		params.Count = BLURT_COUNT
	}

	blurts, err := h.db.GetBlurts(params.Offset, BLURT_COUNT)
	if err != nil {
		slog.Info("error getting blurts", "error", err)
		return err
	}

	return c.JSON(http.StatusOK, blurts)

}
