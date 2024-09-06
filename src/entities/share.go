package entities

import (
	"time"

	"github.com/google/uuid"
)

type ShareEntity struct {
	ShareId     uuid.UUID `gorm:"column:share_id;type:uuid;default:gen_random_uuid();primaryKey"`
	StartDate   time.Time `gorm:"column:start_date;"`
	EndDate     time.Time `gorm:"column:end_date;"`
	ExpiredDate time.Time `gorm:"column:expired_date;"`
	UserShareId uuid.UUID `gorm:"column:user_share_id;"`
}

func (ShareEntity) TableName() string {
	return "share"
}

// `gorm:"foreignKey:SpendingTypeId;references:SpendingTypeId"`
