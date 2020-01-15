package model

import "time"

type User struct {
	UserId string
	UserName string
	Password string
	Email string
	Role string
	CreateAt time.Time
	UpdateAt time.Time
	Token string
}