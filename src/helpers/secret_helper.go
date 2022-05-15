package helpers

import 	"golang.org/x/crypto/bcrypt"

func SecretCodeVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}