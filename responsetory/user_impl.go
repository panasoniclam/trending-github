package responsetory

import (
	"context"
	"github.com/panasoniclam/trending-github/model"
)

type User_repository interface {
	 SaveUser(context context.Context, user model.User) (model.User, error)
}