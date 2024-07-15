package repositories

import (
	"money-service/entities"
	"money-service/interfaces"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepositoryGorm struct {
	db *gorm.DB
}

func UserRepositoryGorm(db *gorm.DB) IUserRepository {
	return &userRepositoryGorm{db: db}
}

func (r *userRepositoryGorm) CreateUser(payload interfaces.UserInsertDb) (*uuid.UUID, error) {
	db := r.db
	newUser := entities.UserEntity{
		UserName:    payload.UserName,
		Email:       payload.Email,
		LastName:    payload.LastName,
		FirstName:   payload.FirstName,
		DisplayName: payload.DisplayName,
		Password:    payload.Password,
		IsActive:    true,
		Profile:     payload.Profile,
	}
	if err := db.Create(newUser).Error; err != nil {
		return nil, err
	}

	return &newUser.UserId, nil
}
func (r *userRepositoryGorm) DeActiveAccount(id string) (bool, error) {
	var result entities.UserEntity
	db := r.db

	if err := db.Model(&result).Update("is_active", false).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *userRepositoryGorm) GetUserById(id uuid.UUID) (*interfaces.UserResultQuery, error) {
	var result entities.UserEntity
	result.UserId = id
	db := r.db

	if err := db.Select("*").First(&result).Where("is_active", true).Error; err != nil {
		return nil, err
	}
	var spendingTypeResults = interfaces.UserResultQuery{
		UserId:      result.UserId,
		UserName:    result.UserName,
		Email:       result.Email,
		LastName:    result.LastName,
		FirstName:   result.FirstName,
		DisplayName: result.DisplayName,
		Profile:     result.Profile,
		CreatedDate: result.CreatedDate,
	}

	return &spendingTypeResults, nil
}
