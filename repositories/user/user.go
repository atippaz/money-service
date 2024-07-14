package repositories

import "money-service/interfaces"

type IUserRepository interface {
	GetUserById(id string) (*interfaces.UserResult, error)
}
