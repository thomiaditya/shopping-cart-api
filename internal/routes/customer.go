package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thomiaditya/shop-api/internal/http/controller"
)

func customerRoutes(source *echo.Group) {
	customerRouteGroup := source.Group("/customer")

	customerController := controller.NewCustomerController()
	customerRouteGroup.GET("", customerController.RegisterCustomer)
}
