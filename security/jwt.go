package security

import (
	"github.com/dgrijalva/jwt-go"
	"myapp/model"
	"time"
)

const SecretKey = "!kHW&kZ35(x)s00!kHW&kZ35(x)s00!kHW&kZ35(x)s00"

func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}
	return result, nil
}
