package handler

import (

	"github.com/labstack/echo"
	"net/http"
	"github.com/panasoniclam/trending-github/Db"
)
type User struct {
   sql *Db.Sql
}
func GetUSer(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "hihihi")
}
type Person struct {
	First_name string `db:"first_name"`
	Last_name string `db:"last_name"`
	Email string
}
