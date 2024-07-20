package services

import (
	user "money-service/src/services/user"
	Hasher "money-service/src/utils/hasher"
	Jwt "money-service/src/utils/jwt"

	"github.com/google/uuid"
)

type AuthService interface {
	Register(payload AuthRegisterInsert) (*uuid.UUID, error)
	Login(credential, password string) (*string, error)
}
type authService struct {
	userService user.UserService
	hasher      Hasher.Hasher
	jwt         Jwt.Jwt
}
type AuthClaims struct {
	Username string    `json:"userName"`
	Email    string    `json:"email"`
	UserId   uuid.UUID `json:"userId"`
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
