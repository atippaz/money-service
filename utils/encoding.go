package utils

import "golang.org/x/crypto/bcrypt"

type IEncodingHelper interface {
	ComparePassword(password, hash string) bool
	Hashing(text string) (*string, error)
}
type encodingHelper struct {
}

func Encoding() IEncodingHelper {
	return &encodingHelper{}
}

func (s *encodingHelper) ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func (s *encodingHelper) Hashing(text string) (*string, error) {
	byteHashing, err := bcrypt.GenerateFromPassword([]byte(text), 12)
	result := string(byteHashing)
	return &result, err
}
