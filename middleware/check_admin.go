package middleware

import (
	"github.com/labstack/echo/v4"
	"myapp/log"
	"myapp/model"
	"myapp/model/req"
	"net/http"
)

func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// get request values
			reqSignIn := req.ReqSignIn{}
			if err := c.Bind(&reqSignIn); err != nil {
				log.Error(err.Error())
				return c.JSON(http.StatusBadRequest, model.Response{
					Status:  http.StatusBadRequest,
					Message: err.Error(),
					Data:    nil,
				})
			}
			// check admin
			if reqSignIn.Email != "admin@gmail.com" {
				return c.JSON(http.StatusBadRequest, model.Response{
					Status:  http.StatusBadRequest,
					Message: "you are not admin",
					Data:    nil,
				})
			}

			return next(c)
		}
	}
}
