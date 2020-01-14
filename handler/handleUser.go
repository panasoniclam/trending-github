package handler

import (
	"github.com/labstack/echo"
	"github.com/panasoniclam/trending-github/model"
	req2 "github.com/panasoniclam/trending-github/model/request"
	"github.com/panasoniclam/trending-github/model/response"
   validator 	"github.com/go-playground/validator/v10"
	"github.com/panasoniclam/trending-github/security"
	 uuid "github.com/google/uuid"
	"net/http"
	"time"
)
type HandleUser struct {

}

func (u *HandleUser)HandleSinUp(context echo.Context)  error {
	req:= req2.RequestSignUp{}
	if err:= context.Bind(&req) ; err != nil {
		  return context.JSON(http.StatusBadRequest, response.Response{
			  StatusCode: http.StatusBadRequest,
			  Message:    err.Error(),
			  Data:       nil,
		  })
	}
	validate := validator.New()
	if err:= validate.Struct(req); err!= nil {
		return  context.JSON(http.StatusBadRequest,response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	hash  := security.HashAndSalt([]byte(req.Password))
	role :=   model.MEMBER.String()

	UserId, err := uuid.NewUUID()
	if err!= nil {
		return context.JSON(http.StatusForbidden,response.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})

	}
   user:= model.User{
	   UserId:   UserId.String(),
	   Email:    req.Email,
	   FullName: req.FullName,
	   Password: hash,
	   Role:     role,
	   CreateAt: time.Time{},
	   UpdateAt: time.Time{},
	   Token:    "",
   }


}