package repositories

import (
	"money-service/entities"
	"money-service/interfaces"

	"gorm.io/gorm"
)

type userRepositoryGorm struct {
	db *gorm.DB
}

func UserRepositoryGorm(db *gorm.DB) IUserRepository {
	return &userRepositoryGorm{db: db}
}

func (r *userRepositoryGorm) GetUserById(id string) (*interfaces.UserResultQuery, error) {
	var result entities.UserEntity
	db := r.db

	if err := db.Select("spending_type_id, name_th, name_en").First(&result).Error; err != nil {
		return nil, err
	}
	var spendingTypeResults = interfaces.UserResultQuery{
		UserId: result.UserId,
	}

	return &spendingTypeResults, nil
}
