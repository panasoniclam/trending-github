package model

import "time"

type (
	User struct {
		UserId string `db:"user_id"`
		FullName string  `db:"full_name"`
		Email string  `db:"email"`
		Password string `db:"password"`
		Role string  `db:"role"`
		CreatedAt time.Time `db:"create_at"`
		UpdatedAt time.Time `db:"update_at"`
		Token string  `db:"tolken"`
	}
)
