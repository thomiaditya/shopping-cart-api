package routes

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type APIRouter struct {
	server *echo.Echo
	api    *echo.Group
}

func NewAPIRouter(server *echo.Echo) *APIRouter {
	return &APIRouter{
		server: server,
		api:    server.Group("/api"),
	}
}

func (router *APIRouter) RegisterRoutes(routes ...func(*echo.Group)) {
	for _, route := range routes {
		route(router.api)
	}
}

func RegisterAPIRoutes(ctx context.Context, server *echo.Echo) {
	// Register the API routes
	apiRouter := NewAPIRouter(server)

	apiRouter.RegisterRoutes(
		customerRoutes,
	)

	// Register the default route
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "It works!")
	})
}
