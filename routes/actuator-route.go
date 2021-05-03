package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetActuatorRoutes(e *echo.Echo) {
	e.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "UP6",
		})
	})
}
