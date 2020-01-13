package route

import (
	"github.com/labstack/echo"
	"github.com/panasoniclam/trending-github/handler"
)

type API struct {
	ECHO *echo.Echo
	HandleUser handler.HandleUser
}

func (api *API)SetupRoute()  {
	api.ECHO.POST("/user/user-sign-up",api.HandleUser.HandleSignIp)
	//api.ECHO.GET("/user/user-sign-in",api.HandleUser.HandleSignIn)
}