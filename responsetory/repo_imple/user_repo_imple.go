package repo_imple

import (
	"context"
	"github.com/lib/pq"
	"github.com/panasoniclam/trending-github/banana"

	"github.com/panasoniclam/trending-github/Db"
	"github.com/panasoniclam/trending-github/model"
	"time"
)

type User_Repo struct {
 	sql *Db.Sql
 }

func (u *User_Repo) SaveUser(context context.Context, user model.User) (model.User,error)  {
	 statement := ` INSERT INTO users(user_id, full_name, email,password,role,create_at,update_at,token)
       VALUES (:user_id,:full_name,:email,:password,:role,:create_at,:update_at,token)
       `
	 user.CreatedAt = time.Now()
	 user.UpdatedAt = time.Now()

	 _,err := u.sql.Db.NamedExecContext(context,statement,user)

	 if err != nil {
         if err, ok := err.(*pq.Error) ; ok  {
         	 if err.Code.Name() == "unique_violation" {
             		return user,banana.UserConfList
			 }
		 }
		 return  user, banana.SignUpFailed
	 }
	 return user, nil
}