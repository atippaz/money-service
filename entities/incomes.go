package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type IncomesEntity struct {
	IncomeId    uuid.UUID       `gorm:"column:income_id;type:uuid;default:gen_random_uuid();primaryKey;"`
	CreatedDate time.Time       `gorm:"column:created_date;not null;default:date_trunc('day'::text, CURRENT_TIMESTAMP)"`
	TagId       uuid.UUID       `gorm:"column:tag_id;not null;"`
	Value       decimal.Decimal `gorm:"column:value;not null;"`
	UserOwner   uuid.UUID       `gorm:"column:user_owner;not null;"`
}

func (IncomesEntity) TableName() string {
	return "incomes"
}

// `gorm:"foreignKey:SpendingTypeId;references:SpendingTypeId"`
