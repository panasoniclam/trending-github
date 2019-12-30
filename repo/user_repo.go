package repo

import  (
	model "github.com/panasoniclam/trending-github/model"
)

type UserRepo interface {
	Select() ([] model.User,error)
	Insert(u  model.User) (error)
}
