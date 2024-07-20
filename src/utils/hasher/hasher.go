package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func NewHasher() Hasher {
	return &hasher{}
}

func (s *hasher) ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err != nil
}
func (s *hasher) Hashing(text *string) (string, error) {
	fmt.Print(text)
	byteHashing, err := bcrypt.GenerateFromPassword([]byte(*text), 12)
	result := string(byteHashing)
	return result, err
}
