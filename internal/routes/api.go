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

func (router *APIRouter) RegisterRoutes(routers ...RouterInterface) {
	for _, route := range routers {
		route.RegisterRoutes(router.api)
	}
}

// Register the API routes in this function
func (api *APIRouter) RegisterAPIRoutes(ctx context.Context) error {
	api.RegisterRoutes(
		NewCustomerRouter(),
	)

	// Register the default route
	api.server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "It works!")
	})

	return nil
}
