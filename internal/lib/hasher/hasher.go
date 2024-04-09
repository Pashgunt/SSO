package hasher

import "golang.org/x/crypto/bcrypt"

func MakeHashPassword(password string) (string, error) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
