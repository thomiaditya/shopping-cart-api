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
	Server *echo.Echo
}

func NewApp() *App {
	return &App{
		Server: echo.New(),
	}
}

func (app *App) Start(ctx context.Context) error {
	// Load the configuration
	err := godotenv.Load()
	if err != nil {
		return err
	}

	// Connect to the database
	db := database.GetDatabaseInstance()
	err = db.Connect(ctx)
	if err != nil {
		return err
	}

	// Register the routes
	apiRouter := routes.NewAPIRouter(app.Server)
	err = apiRouter.RegisterAPIRoutes(ctx)
	if err != nil {
		return err
	}

	// Start the server
	port := util.GetEnv("PORT", "8080")
	portFormatted := fmt.Sprintf(":%s", port)
	return app.Server.Start(portFormatted)
}
