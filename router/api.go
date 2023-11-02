package router

import (
	"myapp/handler"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo *echo.Echo
	Userhandle handler.UserHandle
}

func (api *API) SetUpRouter() {
	api.Echo.POST("/sign-in", api.Userhandle.HandleSignIn)
	api.Echo.POST("/sign-up", api.Userhandle.HandleSignUp)
}