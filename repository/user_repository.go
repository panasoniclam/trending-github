package repository

import (
	"context"
	"github.com/panasoniclam/trending-github/model"
)

type User_Repo interface {
	SaveUser(context context.Context, user model.User) (model.User , error)
}
