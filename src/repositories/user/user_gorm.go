package repositories

import (
	"money-service/src/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepositoryGorm struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryGorm{db: db}
}

func (r *userRepositoryGorm) CreateUser(payload UserInsertDb) (*uuid.UUID, error) {
	db := r.db
	newUser := entities.UserEntity{
		UserId:      uuid.New(),
		UserName:    payload.UserName,
		Email:       payload.Email,
		LastName:    payload.LastName,
		FirstName:   payload.FirstName,
		DisplayName: payload.DisplayName,
		Password:    payload.Password,
		IsActive:    true,
		Profile:     payload.Profile,
	}
	if err := db.Create(&newUser).Error; err != nil {
		return nil, err
	}

	return &newUser.UserId, nil
}
func (r *userRepositoryGorm) DeActiveAccount(userId uuid.UUID) (bool, error) {
	var result entities.UserEntity
	db := r.db

	if err := db.Model(&result).Update("is_active", false).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *userRepositoryGorm) GetUserById(id uuid.UUID) (*UserResultQuery, error) {
	var result entities.UserEntity
	result.UserId = id
	db := r.db

	if err := db.First(&result).Where("is_active", true).Error; err != nil {
		return nil, err
	}
	var spendingTypeResults = UserResultQuery{
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

func (r *userRepositoryGorm) GetUserByCredential(credential string) (*UserResultQuery, error) {
	var result entities.UserEntity
	db := r.db

	if err := db.Where("(email = ? OR user_name = ?) AND is_active = ?", credential, credential, true).First(&result).Error; err != nil {
		return nil, err
	}

	var spendingTypeResults = UserResultQuery{
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
