package main

import (
	"github.com/labstack/echo"
	"github.com/panasoniclam/trending-github/Db"
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

	e.Logger.Fatal(e.Start(":1323"))


}

