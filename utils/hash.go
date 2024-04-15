package utils

import (
	"github.com/hkm007/gopherly/utils/constants"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), constants.HASH_COST)
	return string(bytes), err
}

func CheckPasswordHash(password, hashedPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
