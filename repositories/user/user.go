package repositories

import (
	"money-service/interfaces"

	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserById(id uuid.UUID) (*interfaces.UserResultQuery, error)
	DeActiveAccount(id string) (bool, error)
	CreateUser(payload interfaces.UserInsertDb) (*uuid.UUID, error)
	GetUserByCredential(credential string) (*interfaces.UserResultQuery, error)
}
