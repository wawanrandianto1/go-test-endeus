package repository

import (
	"endeus/wawan/pkg/endeus"
	"endeus/wawan/pkg/model"

	"gorm.io/gorm"
)

type resepRepo struct {
	db *gorm.DB
}

func NewResep(db *gorm.DB) endeus.ResepRepository {
	return &resepRepo{
		db,
	}
}

func (r *resepRepo) GetAll(str string) ([]model.Resep, error) {
	var data []model.Resep
	if str == "" {
		err := r.db.Find(&data).Error
		return data, err
	}

	queryLike := "%" + str + "%"
	err := r.db.Where("judul LIKE ?", queryLike).Or("deskripsi LIKE ?", queryLike).Find(&data).Error
	return data, err
}

func (r *resepRepo) GetAllByCategoryID(categoryID uint) ([]model.Resep, error) {
	var data []model.Resep
	err := r.db.Where("category_id", categoryID).Find(&data).Error
	return data, err
}

func (r *resepRepo) GetByID(id uint) (model.Resep, error) {
	var data model.Resep
	err := r.db.First(&data, id).Error
	return data, err
}

func (r *resepRepo) Create(param endeus.ResepParam) error {
	data := model.Resep{
		Judul:      param.Judul,
		Deskripsi:  param.Deskripsi,
		CategoryID: param.CategoryID,
	}
	err := r.db.Create(&data).Error
	return err
}

func (r *resepRepo) Update(id uint, param endeus.ResepParam) error {
	data := model.Resep{
		ID:         id,
		Judul:      param.Judul,
		Deskripsi:  param.Deskripsi,
		CategoryID: param.CategoryID,
	}
	err := r.db.Save(&data).Error
	return err
}
