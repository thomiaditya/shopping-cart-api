package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thomiaditya/shop-api/internal/http/controller"
)

type CustomerRouter struct{}

func NewCustomerRouter() RouterInterface {
	return &CustomerRouter{}
}

func (router *CustomerRouter) RegisterRoutes(api *echo.Group) error {
	customerRoutes := api.Group("/customer")

	customerController := controller.NewCustomerController()
	customerRoutes.GET("", customerController.RegisterCustomer)
	return nil
}
