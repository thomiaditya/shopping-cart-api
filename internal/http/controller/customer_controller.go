package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomerControllerImplement struct{}

func NewCustomerController() CustomerController {
	return &CustomerControllerImplement{}
}

func (controller *CustomerControllerImplement) RegisterCustomer(echoContext echo.Context) error {
	return echoContext.String(http.StatusOK, "Register Customer")
}
