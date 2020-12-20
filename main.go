package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.GET("/", func(c echo.Context) error {
		return c.String(getHelloWorld())
	})
	e.Logger.Fatal(e.Start(":8080"))
}

func getHelloWorld() (int, string) {
	return 200, "Hello, World!"
}
