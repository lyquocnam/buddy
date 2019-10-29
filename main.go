package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	r := echo.New()
	r.GET("/" , func(c echo.Context) error {
		return c.JSON(200, "Hello world")
	})
	r.Logger.Fatal(r.Start(":3000"))
}
