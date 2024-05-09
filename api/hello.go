package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (controller *Controller) SayHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
