package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (controller *Controller) GetVersion(c echo.Context) error {
	return c.String(http.StatusOK, controller.version)
}
