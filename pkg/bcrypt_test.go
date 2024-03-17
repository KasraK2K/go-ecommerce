package pkg

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "1234567"
	hash, err := HashPassword(password)

	t.Run("HashPassword", func(t *testing.T) {
		if err != nil {
			t.Errorf("error on hash password: %v", err)
		}

		if !(len(hash) > 0) {
			t.Errorf("hash %v is not valid", hash)
		}
	})
}

func TestComparePassword(t *testing.T) {
	password := "1234567"
	hash, _ := HashPassword(password)
	compareHash := ComparePassword(hash, password)

	wrongPassword := "abcdefg"
	compareWrongHash := ComparePassword(hash, wrongPassword)

	t.Run("TestComparePassword", func(t *testing.T) {
		if !compareHash {
			t.Errorf("hash password is not valid")
		}

		if compareWrongHash {
			t.Errorf("compare wrong hash not be valid")
		}
	})
}

func TestRandomHashedPassword(t *testing.T) {
	t.Run("TestRandomHashedPassword", func(t *testing.T) {
		passwordLength := 7
		password, hash, err := RandomHashedPassword(7)

		if err != nil {
			t.Errorf("random hash password has error: %v", err)
		}

		if len(password) != passwordLength {
			t.Errorf("generated password length is not true")
		}

		if !(len(hash) > 0) {
			t.Errorf("generated hash length is not true")
		}
	})
}
