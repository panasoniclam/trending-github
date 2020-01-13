package handler

import (

	"github.com/labstack/echo"
	"github.com/panasoniclam/trending-github/log"
	"github.com/panasoniclam/trending-github/model"
	"github.com/panasoniclam/trending-github/secure"

	"net/http"
   validator  "github.com/go-playground/validator/v10"
	req2 "github.com/panasoniclam/trending-github/model/rep"
	 uuid "github.com/google/uuid"
	"time"
	"github.com/panasoniclam/trending-github/responsetory/repo_imple"
)

type (
	HandleUser struct {
        UserRepo repo_imple.User_Repo
	}
)

func (u *HandleUser)HandleSignIp(context echo.Context) error  {
	req:= req2.RequestSignUp{}
	if err:= context.Bind(&req); err!= nil {
		log.Error(err.Error())
		return  context.JSON(http.StatusBadRequest,model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
    validate := validator.New()
    role :=  model.MEMBER.String()
    hash := secure.HashAndSalt([]byte(req.Password))
    if err:= validate.Struct(req) ; err!= nil {
    	return context.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
    userId, err := uuid.NewUUID()
	if err!= nil{
		return  context.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}
    user:= model.User{
		UserId:    userId.String(),
		FullName:  req.FullName,
		Email:     req.Email,
		Password:  hash,
		Role:      role,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Token:     "",
	}
	user, err = u.UserRepo.SaveUser(context.Request().Context(), user)
	if err!= nil {
		return  context.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	token, err := secure.GenToken(user)

	if err!= nil {
		return context.JSON(http.StatusInternalServerError,model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}


	user.Token = token

	return  context.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "sữ lý thành công",
		Data:       nil,
	})




}