// This file contains the routes interface
package routes

import (
	"github.com/labstack/echo/v4"
)

type RouterInterface interface {
	RegisterRoutes(api *echo.Group) error
}
