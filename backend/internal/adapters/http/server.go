// Package http has all the configuration for the echo server
package http

import (
	"net/http"

	"reservations-backend/internal/adapters/httperr"

	"github.com/labstack/echo/v4"
)

func GetServerInstance() *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = httperr.HTTPErrorHandler

	e.GET("/health", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, "server is live!")
	})

	return e
}
