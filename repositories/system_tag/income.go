package repositories

import "money-service/interfaces"

type ISystemTagRepository interface {
	GetUserById(id string) (*interfaces.UserResult, error)
}
