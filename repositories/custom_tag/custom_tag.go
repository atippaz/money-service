package repositories

import (
	"money-service/interfaces"

	"github.com/google/uuid"
)

type CustomTagRepository interface {
	GetCustomTagsByUser(ownerId string) (*[]interfaces.CustomTagResultQuery, error)
	CreateCustomTag(userOwner uuid.UUID, payload interfaces.CustomTagInsertDB) (*uuid.UUID, error)
}
