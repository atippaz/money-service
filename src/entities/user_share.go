package entities

import (
	"github.com/google/uuid"
)

type UserShareEntity struct {
	UserShareId uuid.UUID `gorm:"column:user_share_id;type:uuid;default:gen_random_uuid();primaryKey"`
	UserId      uuid.UUID `gorm:"column:user_id;type:uuid"`
	ShareId     uuid.UUID `gorm:"column:share_id;type:uuid;"`
}

func (UserShareEntity) TableName() string {
	return "userShareEntity"
}

// `gorm:"foreignKey:SpendingTypeId;references:SpendingTypeId"`
