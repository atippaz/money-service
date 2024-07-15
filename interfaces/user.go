package interfaces

import (
	"time"

	"github.com/google/uuid"
)

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

type UserResultResponse struct {
	UserId      uuid.UUID `json:"userId"`
	Email       string    `json:"email"`
	UserName    string    `json:"userName"`
	LastName    string    `json:"lastName"`
	FirstName   string    `json:"firstName"`
	DisplayName string    `json:"displayName"`
	Profile     string    `json:"profile"`
	CreatedDate time.Time `json:"createdDate"`
}
