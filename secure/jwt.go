package secure

import (
	 "github.com/panasoniclam/trending-github/model"
  	"github.com/dgrijalva/jwt-go"
	"time"
)

const JWT_KEY = "hhhgfdshgfhsdgfshjgfshjdgf"

func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClais{
		UserId: user.UserId,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWT_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

