package api

import (
	"dorobo/web/requests"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func (controller *Controller) Upgrade(c echo.Context) error {
	var req requests.UpgradeRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("unknown request format: %v", err))
	}

	binary, err := os.Open(req.Path)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("failed to open binary: %v", err))
	}

	go func() {
		defer binary.Close()
		if err := controller.pm.Upgrade(binary); err != nil {
			log.Printf("failed to upgrade server: %v", err)
		}
	}()
	return c.NoContent(http.StatusNoContent)
}
