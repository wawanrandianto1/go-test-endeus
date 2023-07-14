package repository

import (
	"endeus/wawan/pkg/endeus"
	"endeus/wawan/pkg/model"

	"github.com/rs/zerolog/log"
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

func (r *resepRepo) FindAll(str string) ([]model.Resep, error) {
	var data []model.Resep
	if str == "" {
		err := r.db.Where("publish = ?", true).Find(&data).Error
		return data, err
	}

	queryLike := "%" + str + "%"
	err := r.db.Where("publish = ? AND judul LIKE ?", true, queryLike).Or("deskripsi LIKE ?", queryLike).Find(&data).Error
	return data, err
}

func (r *resepRepo) FindAllByCategoryID(categoryID uint) ([]model.Resep, error) {
	var data []model.Resep
	err := r.db.Where("category_id", categoryID).Find(&data).Error
	return data, err
}

func (r *resepRepo) FindByID(id uint) (model.Resep, error) {
	var data model.Resep
	err := r.db.Model(&model.Resep{}).Preload("Bahan").Preload("CaraBuat").First(&data, id).Error
	return data, err
}

func (r *resepRepo) Create(param endeus.ResepParam) error {
	data := model.Resep{
		Judul:      param.Judul,
		Deskripsi:  param.Deskripsi,
		CategoryID: param.CategoryID,
		VideoUrl:   param.VideoUrl,
		Publish:    true,
		Bahan: model.Bahan{
			Porsi:     param.Porsi,
			Deskripsi: param.DeskripsiBahan,
		},
		CaraBuat: model.CaraBuat{
			LamaWaktu: param.LamaWaktu,
			Tips:      param.Tips,
			Deskripsi: param.DeskripsiPembuatan,
		},
	}
	err := r.db.Create(&data).Error
	return err
}

func (r *resepRepo) Unpublish(id uint) error {
	return r.db.Model(&model.Resep{ID: id}).Update("publish", false).Error
}

func (r *resepRepo) Delete(id uint) error {
	trans := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trans.Rollback()
			log.Error().Msg("transaction recover")
		}
	}()

	if err := trans.Error; err != nil {
		return err
	}

	if err := trans.Delete(&model.Bahan{}, "resep_id = ?", id).Error; err != nil {
		trans.Rollback()
		log.Error().Err(err).Msg("transaction bahan delete")
		return err
	}

	if err := trans.Delete(&model.CaraBuat{}, "resep_id = ?", id).Error; err != nil {
		trans.Rollback()
		log.Error().Err(err).Msg("transaction carabuat delete")
		return err
	}

	if err := trans.Delete(&model.Resep{}, id).Error; err != nil {
		trans.Rollback()
		log.Error().Err(err).Msg("transaction resep delete")
		return err
	}
	return trans.Commit().Error
}

func (r *resepRepo) Update(id uint, param endeus.ResepParam) error {
	trans := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trans.Rollback()
			log.Error().Msg("transaction recover")
		}
	}()

	if err := trans.Error; err != nil {
		return err
	}

	data := model.Resep{
		ID:         id,
		Judul:      param.Judul,
		Deskripsi:  param.Deskripsi,
		CategoryID: param.CategoryID,
		VideoUrl:   param.VideoUrl,
	}
	if err := trans.Omit("created_at").Updates(&data).Error; err != nil {
		trans.Rollback()
		log.Error().Err(err).Msg("transaction resep update")
		return err
	}

	bahan := model.Bahan{
		Porsi:     param.Porsi,
		Deskripsi: param.DeskripsiBahan,
	}
	err := trans.Model(&model.Bahan{}).Where("resep_id = ?", id).Omit("created_at").Updates(&bahan).Error
	if err != nil {
		trans.Rollback()
		log.Error().Err(err).Msg("transaction bahan update")
		return err
	}

	carabuat := model.CaraBuat{
		LamaWaktu: param.LamaWaktu,
		Tips:      param.Tips,
		Deskripsi: param.DeskripsiPembuatan,
	}
	err = trans.Model(&model.CaraBuat{}).Where("resep_id = ?", id).Omit("created_at").Updates(&carabuat).Error
	if err != nil {
		trans.Rollback()
		log.Error().Err(err).Msg("transaction carabuat update")
		return err
	}
	return trans.Commit().Error
}
