// Package httperr provides custom HTTP error handler
package httperr

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomErr struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HTTPStatus uint   `json:"-"`
}

type HTTPErr struct {
	CustomErr CustomErr `json:"error"`
}

func (e HTTPErr) Error() string {
	return fmt.Sprintf("HTTP Status: %d Error Code: %s Message: %s", e.CustomErr.HTTPStatus, e.CustomErr.Code, e.CustomErr.Message)
}

func HTTPErrorHandler(err error, c echo.Context) error {
	customError, ok := err.(HTTPErr)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "internal server error, check logs"})
	}

	return c.JSON(int(customError.CustomErr.HTTPStatus), err)
}
