package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/jeffreycharters/blurt/internal/blurthub"
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

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	h := handlers.New(db)

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "frontend/build",
		Filesystem: http.FS(webAssets),
	}))

	e.GET("/ws", h.WebsocketHandler)
	go blurthub.Run()

	api := e.Group("/api")

	users := api.Group("/users")
	users.POST("/:username", h.NewUserHandler)
	users.GET("/usercount", h.UserCountHandler)

	blurts := api.Group("/blurts")
	blurts.GET("/", h.GetBlurtsHandler)
	blurts.POST("/", h.NewBlurtHandler)

	liks := api.Group("/liks")
	liks.POST("/", h.NewLikHandler)

	log.Fatal(e.Start(":3320"))

}
