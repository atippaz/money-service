package services

import (
	userRepositories "money-service/src/repositories/user"
	Hasher "money-service/src/utils/hasher"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	GetLoginDataByCredential(credential string) (*UserLoginInfo, error)
	DeActiveAccount(userId uuid.UUID) (bool, error)
	CreateUser(payload UserResult) (*uuid.UUID, error)
	GetUserById(userId uuid.UUID) (*UserInfo, error)
}
type userService struct {
	repo   userRepositories.UserRepository
	encode Hasher.Hasher
}
type UserResult struct {
	UserName    string
	Email       string
	LastName    string
	FirstName   string
	DisplayName string
	Password    string
	Profile     string
}
type UserInfo struct {
	UserId      uuid.UUID
	Email       string
	UserName    string
	LastName    string
	FirstName   string
	DisplayName string
	Profile     string
	CreatedDate time.Time
}
type UserLoginInfo struct {
	UserId   uuid.UUID
	Email    string
	UserName string
	Password string
}
