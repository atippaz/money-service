package repositories

import (
	"fmt"
	"money-service/entities"
	"money-service/interfaces"

	"gorm.io/gorm"
)

type systemTagRepositoryGorm struct {
	db *gorm.DB
}

func SystemTagRepositoryGorm(db *gorm.DB) ISystemTagRepository {
	return &systemTagRepositoryGorm{db: db}
}

func (r *systemTagRepositoryGorm) GetUserById(id string) (*interfaces.UserResult, error) {
	var result entities.UserEntity
	db := r.db

	if err := db.Select("spending_type_id, name_th, name_en").First(&result).Error; err != nil {
		return nil, err
	}
	fmt.Print(result)
	var spendingTypeResults = interfaces.UserResult{
		UserId: result.UserId,
	}

	return &spendingTypeResults, nil
}
