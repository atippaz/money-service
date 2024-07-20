package utils

type Hasher interface {
	ComparePassword(password, hash string) bool
	Hashing(text *string) (string, error)
}
type hasher struct {
}
