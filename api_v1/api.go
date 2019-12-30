package api_v1

import (
	"github.com/labstack/echo"
	"net/http"
)
type (
	User struct {
		 Username string  `json:"hoten"`
		 Email string	`json:"diachi"`
		 Id string
	}
)

func getUser(c echo.Context) error  {
	user := User{
		Username: "lamnn8",
		Email:    "lamnnn8@fpt.com.vn",
		Id:"1",
	}
	id:= c.Param("id")
	if id== user.Id {
		return  c.JSON(http.StatusOK,user)
	}
	return  c.JSON(http.StatusBadRequest,"not dpunt")
}
func Save(c echo.Context) error  {
	user :=  User{
		Username: c.FormValue("username"),
		Email:    c.FormValue("email"),
	}
	
	return  c.JSON(http.StatusOK,user)
}