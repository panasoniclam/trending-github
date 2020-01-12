package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/panasoniclam/trending-github/model"
	req2 "github.com/panasoniclam/trending-github/model/rep"
	"net/http"
)

type (
	HandleUser struct {

	}
)

func (u *HandleUser)HanldeSignIn(context echo.Context) error  {


	return context.JSON(http.StatusOK,echo.Map{
		"Email":"lamnn8@fpt.com",
		"FullName":"nguyen ngoc lam",
	})
}
func (u *HandleUser)HandleSignUp(context echo.Context) error {


	req:= req2.RequestSignUp{}
	if err:= context.Bind(&req) ; err != nil {
		return context.JSON(http.StatusBadRequest,model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	validate := validator.New()
	if err:= validate.Struct(req) ;err!= nil {
		return context.JSON(http.StatusBadRequest,model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	type User struct {
		Email string  `json:"email"`
		FullName string `json:"full"`
	}
	user := User{
		Email:    "lamnn",
		FullName: "lam nguyen ngoc",
	}
	return context.JSON(http.StatusOK,user)
}