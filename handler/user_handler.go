package handler

import (
	validator "github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"myapp/log"
	"myapp/model"
	req "myapp/model/req"
	"myapp/repository"
	"myapp/security"
	"net/http"
)

type UserHandle struct {
	UserRepo repository.UserRepo
}

func (userHandle *UserHandle) HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"name": "Cong Dat",
	})
}

func (userHandle *UserHandle) HandleSignUp(c echo.Context) error {
	req := req.ReqSignUp{}
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

	user := model.User{
		UserId:   userId.String(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: hash,
		Role:     role,
		Token:    "",
	}
	user, err = userHandle.UserRepo.SaveUser(c.Request().Context(), user)

	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			Status:  http.StatusConflict,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "user is added into DB",
		Data:    user,
	})
}
