package router

import (
	"github.com/labstack/echo/v4"
	"myapp/handler"
	mymiddleware "myapp/middleware"
)

type API struct {
	Echo       *echo.Echo
	Userhandle handler.UserHandle
}

func (api *API) SetUpRouter() {
	api.Echo.POST("/sign-in", api.Userhandle.HandleSignIn, mymiddleware.IsAdmin())
	api.Echo.POST("/sign-up", api.Userhandle.HandleSignUp)
}
