package model

import "github.com/dgrijalva/jwt-go"

type JwtCustomClais struct {
	UserId string
	Role string
	jwt.StandardClaims
}
