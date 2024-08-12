package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string,error) {
	result,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("Fail to hash password: %w",err)
	}

	return string(result),nil
}

func CheckPassword(password string,hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword),[]byte(password))
}