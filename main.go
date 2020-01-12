package main

import (
	"github.com/labstack/echo"
	"github.com/panasoniclam/trending-github/Db"
	"github.com/panasoniclam/trending-github/handler"
	"github.com/panasoniclam/trending-github/route"
)


func main()  {
	Db:= Db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "3414pic2ck2pi",
		DbName:   "postgres",
	}
	Db.Connection()
	defer Db.Close()
	

	e:= echo.New()
	api := route.API{
		ECHO:       e,
		HandleUser: handler.HandleUser{},
	}
	api.SetupRoute()
	e.Logger.Fatal(e.Start(":1323"))


}

