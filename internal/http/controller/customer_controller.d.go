package controller

import (
	"github.com/labstack/echo/v4"
)

type CustomerController interface {
	RegisterCustomer(echoContext echo.Context) error
}
