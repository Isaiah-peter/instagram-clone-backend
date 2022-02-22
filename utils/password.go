package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Hashpassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Panic(err)
	}
	return string(hash), err
}
