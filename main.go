package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	middleware.Static("")

	e.Get("/api", func(c echo.Context) error {
		return c.String(200, "yo!")
	})

	e.Run(standard.WithTLS(":1221", "crt/server.crt", "crt/server.key"))
}
