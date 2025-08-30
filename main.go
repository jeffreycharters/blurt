package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/jeffreycharters/blurt/internal/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed public/*
var webAssets embed.FS

func main() {
	db := db.Init()
	defer db.Close()

	app := echo.New()

	app.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "public",
		Filesystem: http.FS(webAssets),
	}))

	app.GET("api/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "users")
	})

	log.Fatal(app.Start(":3320"))

}
