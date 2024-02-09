package utils

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"os"
	"strconv"
)

func HashPassword(password string) string {
	salt, _ := strconv.Atoi(os.Getenv("SALT"))
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes)
}

func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var characters_password = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789*${!:;?+-&}")

func GeneratePassword(size int) string {
	b := make([]rune, size)
	for i := range b {
		b[i] = characters_password[rand.Intn(len(characters_password))]
	}
	return string(b)
}
