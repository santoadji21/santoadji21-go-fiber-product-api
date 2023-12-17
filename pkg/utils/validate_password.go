package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(submittedPassword, storedPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(submittedPassword))
    return err == nil
}
