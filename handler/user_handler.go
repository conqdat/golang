package handler

import (
	validator "github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"myapp/log"
	"myapp/model"
	req2 "myapp/model/req"
	"myapp/security"
	"net/http"
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
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	hash := security.HashAndSalt([]byte(req.Password))
	role := model.MEMBER.String()

	userId, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusForbidden, model.Response{
			Status:  http.StatusBadRequest,
			Message: "Fail to gen UUID",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, currentUser)
}
