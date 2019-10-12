package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func CheckHealth(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
