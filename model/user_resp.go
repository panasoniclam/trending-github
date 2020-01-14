package model

import "time"

type User struct {
	UserId string
	Email string
	FullName string
	Password string
	Role string
	CreateAt time.Time
	UpdateAt time.Time
	Token string
}