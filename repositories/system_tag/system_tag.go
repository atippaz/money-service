package repositories

import "money-service/interfaces"

type SystemTagRepository interface {
	GetAllSystemTags() (*[]interfaces.SystemTagResultQuery, error)
}
