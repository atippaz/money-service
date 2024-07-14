package entities

import (
	"github.com/google/uuid"
)

type SystemTagEntity struct {
	TagId          uuid.UUID `gorm:"column:tag_id;type:uuid;default:gen_random_uuid();primaryKey;"`
	NameTh         string    `gorm:"column:name_th"`
	NameEn         string    `gorm:"column:name_en"`
	IsActive       bool      `gorm:"column:is_active;not null;default:true;"`
	SpendingTypeId uuid.UUID `gorm:"column:spending_type_id;not null;"`
}

func (SystemTagEntity) TableName() string {
	return "system_tags"
}

// `gorm:"foreignKey:SpendingTypeId;references:SpendingTypeId"`
