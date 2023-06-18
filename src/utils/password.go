package utils

import bycrypt "golang.org/x/crypto/bcrypt"

func HashPassword(value string) ([]byte, error) {
	return bycrypt.GenerateFromPassword([]byte(value), bycrypt.DefaultCost)

}

func CheckPasswordHash(password, hash string) error {
	err := bycrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err
}
