package pkg

import (
	"math/rand"
	"time"
)

func ReverseString(str string) string {
	var revStr string

	for _, s := range str {
		revStr = string(s) + revStr
	}

	return revStr
}

func RandomString(length int) string {
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[randSource.Intn(len(charset))]
	}
	return string(b)
}
