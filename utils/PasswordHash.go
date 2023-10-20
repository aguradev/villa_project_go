package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	hashing, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", nil
	}

	return string(hashing), nil

}

func ComparePasswordHashing(password, hash string) (bool, error) {

	compareException := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if compareException != nil {
		return false, compareException
	}

	return true, nil

}
