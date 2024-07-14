package repositories

import "money-service/interfaces"

type ICustomTagRepository interface {
	GetUserById(id string) (*interfaces.UserResult, error)
}
