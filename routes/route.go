package routes

import (
	"demo_echo/handler"

	"github.com/labstack/echo"
)

func Public(e *echo.Echo) {
	g := e.Group("/api/v1/public")
	g.GET("/health", handler.CheckHealth)
	g.GET("/student_list", handler.GetStudentList)
}
