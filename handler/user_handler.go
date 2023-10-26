package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"name": "Cong Dat",
	})
}

func HandleSignUp(c echo.Context) error {

	type User struct {
		YourName string `json:"your_name"`
		Gmail    string
	}

	currentUser := User{
		YourName: "Cong DAt",
		Gmail:    "congdat@gmail.com",
	}

	return c.JSON(http.StatusOK, currentUser)
}
