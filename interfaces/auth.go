package interfaces

import "github.com/google/uuid"

type AuthRegisterRequest struct {
	UserName    string `json:"userName"`
	Email       string `json:"email"`
	LastName    string `json:"lastName"`
	FirstName   string `json:"firstName"`
	DisplayName string `json:"displayName"`
	Password    string `json:"password"`
	Profile     string `json:"profile"`
}

type AuthLoginRequest struct {
	Credential string `json:"credential"`
	Password   string `json:"password"`
}
type AuthRegisterInsert struct {
	UserName    string
	Email       string
	LastName    string
	FirstName   string
	DisplayName string
	Password    string
	Profile     string
}

type AuthClaims struct {
	Username string    `json:"userName"`
	Email    string    `json:"email"`
	UserId   uuid.UUID `json:"userId"`
	// jwt.RegisteredClaims
}
