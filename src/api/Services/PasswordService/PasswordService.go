package PasswordService

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	salt = 10
)

func IsValid(hash string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func Encrypt(password string) (string, error) {
	str, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(str), err
}
