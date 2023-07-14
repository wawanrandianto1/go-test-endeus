package endeus

import (
	"endeus/wawan/pkg/model"
)

//go:generate mockgen -source=repo.go -package=mocks -destination=../../internal/mocks/repo.go

type CategoryRepository interface {
	FindAll() ([]model.Category, error)
}

type ResepRepository interface {
	FindAll(str string) ([]model.Resep, error)
	FindAllByCategoryID(categoryID uint) ([]model.Resep, error)
	FindByID(resepID uint) (model.Resep, error)
	Create(param ResepParam) error
	Unpublish(id uint) error
	Delete(id uint) error
	Update(id uint, param ResepParam) error
}

type ResepParam struct {
	CategoryID         uint   `json:"category_id" validate:"required"`
	Judul              string `json:"judul" validate:"required"`
	Deskripsi          string `json:"deskripsi" validate:"required"`
	VideoUrl           string `json:"video_url" validate:"required"`
	Porsi              uint   `json:"porsi" validate:"required"`
	DeskripsiBahan     string `json:"deskripsi_bahan" validate:"required"`
	LamaWaktu          uint   `json:"lama_waktu" validate:"required"`
	Tips               string `json:"tips"`
	DeskripsiPembuatan string `json:"deskripsi_pembuatan" validate:"required"`
}
