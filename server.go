package main

import (
	"myapp/db"
	"myapp/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "password",
		DbName:   "postgres",
	}

	sql.Connect()
	defer sql.Close()

	e := echo.New()

	e.GET("/", home)
	e.GET("/sign-in", handler.HandleSignIn)
	e.GET("/sign-up", handler.HandleSignUp)

	e.Logger.Fatal(e.Start(":3000"))
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
