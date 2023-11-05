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
	reqSignIn := req.ReqSignIn{}
	if err := c.Bind(&reqSignIn); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}
	validate := validator.New()
	if err := validate.Struct(reqSignIn); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	user, err := userHandle.UserRepo.CheckLogin(c.Request().Context(), reqSignIn)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusUnauthorized, model.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// check pass
	isTheSame := security.ComparePasswords(user.Password, []byte(reqSignIn.Password))
	if !isTheSame {
		log.Error(err.Error())
		return c.JSON(http.StatusUnauthorized, model.Response{
			Status:  http.StatusBadRequest,
			Message: "Fail to login",
			Data:    nil,
		})
	}

	// gen token
	token, err := security.GenToken(user)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusUnauthorized, model.Response{
			Status:  http.StatusBadRequest,
			Message: "Fail to generate token",
			Data:    nil,
		})
	}
	user.Token = token

	return c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "yes, this is a valid user",
		Data:    user,
	})
}

func (userHandle *UserHandle) HandleSignUp(c echo.Context) error {
	reqSignUp := req.ReqSignUp{}
	if err := c.Bind(&reqSignUp); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	validate := validator.New()
	if err := validate.Struct(reqSignUp); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	hash := security.HashAndSalt([]byte(reqSignUp.Password))
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
		FullName: reqSignUp.FullName,
		Email:    reqSignUp.Email,
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

	// gen token
	token, err := security.GenToken(user)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusUnauthorized, model.Response{
			Status:  http.StatusBadRequest,
			Message: "Fail to generate token",
			Data:    nil,
		})
	}
	user.Token = token

	return c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "user is added into DB",
		Data:    user,
	})

}

func (userHandle *UserHandle) HandleGetProfile(c echo.Context) error {
	return nil
}
