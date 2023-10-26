package main

import (
	"myapp/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", home)
	e.GET("/sign-in", handler.HandleSignIn)
	e.GET("/sign-up", handler.HandleSignUp)

	e.Logger.Fatal(e.Start(":3000"))
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
