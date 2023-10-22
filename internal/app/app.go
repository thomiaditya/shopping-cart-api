package app

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/thomiaditya/shop-api/internal/database"
	"github.com/thomiaditya/shop-api/internal/routes"
	"github.com/thomiaditya/shop-api/util"
)

type App struct {
	server *echo.Echo
	db     database.DatabaseInterface
}

func NewApp() *App {
	return &App{
		server: echo.New(),
	}
}

func (app *App) Start(ctx context.Context) error {
	// Load the configuration
	err := godotenv.Load()
	if err != nil {
		return err
	}

	// Connect to the database
	app.db = database.NewPostgresDatabaseFromConfig()
	err = app.db.Connect(ctx)
	if err != nil {
		return err
	}

	// Register the routes
	routes.RegisterAPIRoutes(ctx, app.server)

	// Start the server
	port := util.GetEnv("PORT", "8080")
	portFormatted := fmt.Sprintf(":%s", port)
	return app.server.Start(portFormatted)
}
