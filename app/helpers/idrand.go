package helpers

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func GenerateID() string {
	length := rand.Intn(8) + 4
	randomPart := GenerateRandomString(length)
	datePart := time.Now().Format("060102") // YYMMDD format
	return randomPart + datePart
}
