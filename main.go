package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/jeffreycharters/blurt/internal/database"
	"github.com/jeffreycharters/blurt/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed frontend/build/*
var webAssets embed.FS

func main() {
	db := database.Init()
	defer db.Close()

	app := echo.New()
	app.Use(middleware.CORS())

	h := handlers.New(db)
	h.Routes = app

	app.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "frontend/build",
		Filesystem: http.FS(webAssets),
	}))

	app.GET("api/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "users")
	})

	app.POST("api/user/:username", h.NewUserHandler)
	app.GET("api/usercount", h.UserCountHandler)

	log.Fatal(app.Start(":3320"))

}
