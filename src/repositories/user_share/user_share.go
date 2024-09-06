package repositories

import "github.com/google/uuid"

type UserShareRepository interface {
	Insert(userOwner uuid.UUID, payload UserShareInsertDB) (*uuid.UUID, error)
	GetAll(userId *uuid.UUID) (*[]UserShareResultQuery, error)
}
type UserShareResultQuery struct {
	UserShareId uuid.UUID
	UserId      uuid.UUID
	ShareId     uuid.UUID
}

type UserShareInsertDB struct {
	UserId uuid.UUID
}
