package entities

import (
	"github.com/google/uuid"
)

type SpendingTypeEntity struct {
	SpendingTypeId uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:spending_type_id"`
	NameTh         string    `gorm:"unique;not null;column:name_th"`
	NameEn         string    `gorm:"unique;not null;column:name_en"`
}

func (SpendingTypeEntity) TableName() string {
	return "spending_types"
}

// `gorm:"foreignKey:SpendingTypeId;references:SpendingTypeId"`
