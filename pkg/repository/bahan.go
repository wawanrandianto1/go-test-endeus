package repository

import (
	"endeus/wawan/pkg/endeus"
	"endeus/wawan/pkg/model"

	"gorm.io/gorm"
)

type bahanRepo struct {
	db *gorm.DB
}

func NewBahan(db *gorm.DB) endeus.BahanRepository {
	return &bahanRepo{
		db,
	}
}

func (r *bahanRepo) GetByResepID(resepID uint) (model.Bahan, error) {
	var data model.Bahan
	err := r.db.Where("resep_id = ?", resepID).First(&data).Error
	return data, err
}
