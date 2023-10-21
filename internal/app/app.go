package app

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/thomiaditya/shop-api/internal/routes"
)

type App struct {
	server *echo.Echo
}

func NewApp() *App {
	return &App{
		server: echo.New(),
	}
}

func (app *App) Start(ctx context.Context) error {
	// Register the routes
	routes.RegisterAPIRoutes(ctx, app.server)

	// Start the server
	return app.server.Start(":8080")
}
