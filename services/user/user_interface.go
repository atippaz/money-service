package services

import (
	userRepositories "money-service/repositories/user"
	Hasher "money-service/utils/hasher"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	GetLoginDataByCredential(credential string) (*UserLoginInfo, error)
	DeActiveAccount(id string) (bool, error)
	CreateUser(payload UserResult) (*uuid.UUID, error)
	GetUserById(id uuid.UUID) (*UserInfo, error)
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
	UserId      uuid.UUID `json:"userId"`
	Email       string    `json:"email"`
	UserName    string    `json:"userName"`
	LastName    string    `json:"lastName"`
	FirstName   string    `json:"firstName"`
	DisplayName string    `json:"displayName"`
	Profile     string    `json:"profile"`
	CreatedDate time.Time `json:"createdDate"`
}
type UserLoginInfo struct {
	UserId   uuid.UUID `json:"userId"`
	Email    string    `json:"email"`
	UserName string    `json:"userName"`
	Password string    `json:"password"`
}
