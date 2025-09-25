package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword(password string) (string, error) {
	passBytes := []byte(password)
	cost := bcrypt.DefaultCost

	hashedPassword, err := bcrypt.GenerateFromPassword(passBytes, cost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func VerifyHashedPassword(password string, hashedPassword string) error {
	passBytes := []byte(password)
	hashedPassBytes := []byte(hashedPassword)

	if err := bcrypt.CompareHashAndPassword(hashedPassBytes, passBytes); err != nil {
		return err
	}

	return nil
}
