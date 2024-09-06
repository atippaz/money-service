package repositories

import (
	"time"

	"github.com/google/uuid"
)

type ShareRepository interface {
	Insert(userOwner uuid.UUID, payload ShareInsertDB) (*uuid.UUID, error)
	GetAll(userId *uuid.UUID) (*[]ShareResultQuery, error)
	GetById(shareId *uuid.UUID) (*[]ShareResultQuery, error)
}
type ShareResultQuery struct {
	ShareId     uuid.UUID
	StartDate   time.Time
	EndDate     time.Time
	ExpiredDate time.Time
	UserShareId uuid.UUID
}

type ShareInsertDB struct {
	StartDate   time.Time
	EndDate     time.Time
	ExpiredDate time.Time
	UserShareId uuid.UUID
}
