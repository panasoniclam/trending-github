package handle

import (
	"github.com/labstack/echo"
	"github.com/panasoniclam/trending-github/model"
	req2 "github.com/panasoniclam/trending-github/model/req"
	"github.com/panasoniclam/trending-github/security"
	"net/http"
	validator "github.com/go-playground/validator/v10"
	 uuid "github.com/google/uuid"
	"time"
)

type HanlderUser struct {
	
}

func (u *HanlderUser)HandleSignUp(context echo.Context) error  {
	req:=  req2.UserRequest{}
	if err := context.Bind(&req) ; err != nil {
		return context.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	validate := validator.New()
	if err:= validate.Struct(req) ; err!= nil {
		return  context.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}



	UserId, err := uuid.NewUUID()
	if err != nil {
		return context.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:  err.Error() ,
			Data:       nil,
		})
	}
	hash  := security.HashAndSalt([]byte(req.Pssword))
	role := model.MEMBER.String()
	user := model.User{
		UserId:   UserId.String(),
		UserName: req.FullName,
		Password: hash,
		Email:    req.Email,
		Role: role,
		CreateAt: time.Time{},
		UpdateAt: time.Time{},
		Token:    "",
	}





}
