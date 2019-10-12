package main

import (
	"fmt"

	"github.com/labstack/echo"

	"demo_echo/config"
	"demo_echo/routes"
)

func main() {
	fmt.Println(config.Config)
	e := echo.New()
	routes.Public(e)

	e.Logger.Fatal(e.Start(":3000"))

}
