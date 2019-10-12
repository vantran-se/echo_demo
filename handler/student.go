package handler

import (
	"net/http"

	"demo_echo/models"

	"github.com/labstack/echo"
)

func GetStudentList(c echo.Context) error {
	return c.JSON(http.StatusOK, models.StudentList)
}
