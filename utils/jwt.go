package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hkm007/gopherly/utils/constants"
)

func GenerateToken(email string, userId int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 4).Unix(),
	})

	return token.SignedString([]byte(constants.JWT_SECRET_KEY))
}
