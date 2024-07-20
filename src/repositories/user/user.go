package repositories

import (
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserById(id uuid.UUID) (*UserResultQuery, error)
	DeActiveAccount(id string) (bool, error)
	CreateUser(payload UserInsertDb) (*uuid.UUID, error)
	GetUserByCredential(credential string) (*UserResultQuery, error)
}
type UserInsertDb struct {
	UserName    string
	Email       string
	LastName    string
	FirstName   string
	DisplayName string
	Password    string
	Profile     string
}

type UserResultQuery struct {
	UserId      uuid.UUID
	UserName    string
	Email       string
	LastName    string
	FirstName   string
	DisplayName string
	Password    string
	IsActive    bool
	Profile     string
	CreatedDate time.Time
}
type UserLoginInfoQuery struct {
	UserId   uuid.UUID `json:"userId"`
	Email    string    `json:"email"`
	UserName string    `json:"userName"`
	Password string    `json:"password"`
}
