package repositories

import "money-service/interfaces"

type ICustomTagRepository interface {
	GetCustomTagsByUser(ownerId string) (*[]interfaces.CustomTagResultQuery, error)
}
