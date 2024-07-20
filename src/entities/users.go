package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserEntity struct {
	UserId      uuid.UUID `gorm:"column:user_id;type:uuid;default:gen_random_uuid();primaryKey"`
	UserName    string    `gorm:"column:user_name;"`
	Email       string    `gorm:"column:email;"`
	LastName    string    `gorm:"column:last_name;"`
	FirstName   string    `gorm:"column:first_name;"`
	DisplayName string    `gorm:"column:display_name;"`
	Password    string    `gorm:"column:password;"`
	IsActive    bool      `gorm:"column:is_active;"`
	Profile     string    `gorm:"column:profile;"`
	CreatedDate time.Time `gorm:"column:created_date;"`
}

func (UserEntity) TableName() string {
	return "users"
}

// `gorm:"foreignKey:SpendingTypeId;references:SpendingTypeId"`
