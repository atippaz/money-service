package entities

import (
	"github.com/google/uuid"
)

type CustomTagEntity struct {
	TagId          uuid.UUID `gorm:"column:tag_id;type:uuid;default:gen_random_uuid();primaryKey;"`
	NameTh         string    `gorm:"column:name_th;"`
	NameEn         string    `gorm:"column:name_en;"`
	IsActive       bool      `gorm:"column:is_active;not null;default:true;"`
	SpendingTypeId uuid.UUID `gorm:"column:spending_type_id"`
	UserOwner      uuid.UUID `gorm:"column:user_owner;not null;"`
}

func (CustomTagEntity) TableName() string {
	return "custom_tags"
}

// `gorm:"foreignKey:SpendingTypeId;references:SpendingTypeId"`
