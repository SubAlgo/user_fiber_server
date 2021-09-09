package token

import (
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

const SecretKey = "secret"

func New(userID uint) (token string, err error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err = claims.SignedString([]byte(SecretKey))
	return
}
