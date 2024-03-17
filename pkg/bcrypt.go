package pkg

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 7)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func RandomHashedPassword(length int) (string, string, error) {
	password := RandomString(length)
	hash, err := HashPassword(password)
	return password, hash, err
}
