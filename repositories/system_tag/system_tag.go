package repositories

import "money-service/interfaces"

type ISystemTagRepository interface {
	GetAllSystemTags() (*[]interfaces.SystemTagResultQuery, error)
}
