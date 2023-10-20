package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type App struct {
	server *echo.Echo
}

func NewApp() *App {
	return &App{
		server: echo.New(),
	}
}

func (app *App) Start() error {
	app.server.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello, world!")
	})

	// Start the server
	return app.server.Start(":8080")
}
