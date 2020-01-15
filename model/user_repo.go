package model

import "time"

type User struct {
	UserId string `db:"user_id"`
	UserName string `db:"user_name"`
	Password string `db:"password"`
	Email string  `db:"email"`
	Role string `db:"role"`
	CreateAt time.Time `db:"create_at"`
	UpdateAt time.Time `db:"update_at"`
	Token string `db:"token"`
}