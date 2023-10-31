package handler

import (
	"myapp/model"
	req2 "myapp/model/req"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandle struct{}

func (userHandle *UserHandle) HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"name": "Cong Dat",
	})
}

func (userHandle *UserHandle) HandleSignUp(c echo.Context) error {
	req := req2.ReqSignUp{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

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
