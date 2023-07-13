package repository

import (
	"endeus/wawan/pkg/endeus"
	"endeus/wawan/pkg/model"

	"gorm.io/gorm"
)

type categoryRepo struct {
	db *gorm.DB
}

func NewCategory(db *gorm.DB) endeus.CategoryRepository {
	return &categoryRepo{
		db,
	}
}

func (r *categoryRepo) GetAll() ([]model.Category, error) {
	var data []model.Category
	err := r.db.Find(&data).Error
	return data, err
}
