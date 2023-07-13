package repository

import (
	"endeus/wawan/pkg/endeus"
	"endeus/wawan/pkg/model"

	"gorm.io/gorm"
)

type caraMasakRepo struct {
	db *gorm.DB
}

func NewMasak(db *gorm.DB) endeus.CaraMasakRepository {
	return &caraMasakRepo{
		db,
	}
}

func (r *caraMasakRepo) GetByResepID(resepID uint) (model.CaraMasak, error) {
	var data model.CaraMasak
	err := r.db.Where("resep_id = ?", resepID).First(&data).Error
	return data, err
}
