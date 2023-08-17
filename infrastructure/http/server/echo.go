package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Server() *echo.Echo {
	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Success",
			"data":    "Welcome to API BRI",
		})
	})

	return app
}
