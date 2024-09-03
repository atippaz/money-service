package repositories

import (
	"money-service/src/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type tagRepositoryGorm struct {
	db *gorm.DB
}

func NewGormTagRepository(db *gorm.DB) TagRepository {
	return &tagRepositoryGorm{db: db}
}
func (r *tagRepositoryGorm) CreateTag(userOwner uuid.UUID, payload TagInsertDB) (*uuid.UUID, error) {
	db := r.db
	newTag := entities.TagEntity{
		NameTh:         payload.NameTh,
		NameEn:         payload.NameEn,
		IsActive:       true,
		SpendingTypeId: payload.SpendingTypeId,
		Owner:          userOwner,
	}
	if err := db.Create(&newTag).Error; err != nil {
		return nil, err
	}
	return &newTag.TagId, nil
}

func (r *tagRepositoryGorm) GetTags(userId *uuid.UUID, hasSystem bool) (*[]TagResultQuery, error) {
	var results []entities.TagEntity
	db := r.db

	if err := db.Where("is_active = ? AND user_owner = ?", hasSystem, userId).Find(&results).Error; err != nil {
		return nil, err
	}
	if userId == nil {
		if err := r.db.Find(&results).Error; err != nil {
			return nil, err
		}
	} else {
		if err := r.db.Where("is_active = ? ANF owner = ?", hasSystem, *userId).Find(&results).Error; err != nil {
			return nil, err
		}
	}
	var TagResults []TagResultQuery
	for _, result := range results {
		TagResults = append(TagResults, TagResultQuery{
			SpendingTypeId: result.SpendingTypeId,
			NameTh:         result.NameTh,
			NameEn:         result.NameEn,
			TagId:          result.TagId,
			IsActive:       result.IsActive,
			Owner:          result.Owner,
		})
	}

	return &TagResults, nil
}
