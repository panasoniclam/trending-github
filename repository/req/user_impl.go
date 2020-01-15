package req

import (
	"context"
	"github.com/lib/pq"
	"github.com/panasoniclam/trending-github/banana"
	"github.com/panasoniclam/trending-github/db"
	"github.com/panasoniclam/trending-github/model"
	"time"
)

type User_Repository struct {
	sql *db.Sql
}

func (u *User_Repository)SaveUser(context context.Context, user model.User)(model.User , error)  {
	  statement := `INSERT INTO users(user_id,user_name,password,email,role,create_at,update_at) VALUES(:user_id,:user_name,:password,:email,:role,:create_at,:update_at)`
	  user.UpdateAt = time.Now()
	  user.CreateAt = time.Now()
	  _, err := u.sql.Db.NamedExecContext(context, statement,user)
	  if err != nil {
	  	if err,ok := err.(*pq.Error) ; ok {
	  		if err.Code.Name() ==  "unique_violation" {
	  			return  user, banana.UserConfigList
			}
		}
		return  user, banana.UserSignUpFaied
	  }
	  return  user, nil
}