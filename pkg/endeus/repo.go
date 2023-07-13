package endeus

import (
	"endeus/wawan/pkg/model"
)

//go:generate mockgen -source=repo.go -package=mocks -destination=../../internal/mocks/repo.go

type CaraMasakRepository interface {
	GetByResepID(resepID uint) (model.CaraMasak, error)
}

type BahanRepository interface {
	GetByResepID(resepID uint) (model.Bahan, error)
}

type CategoryRepository interface {
	GetAll() ([]model.Category, error)
}

type ResepRepository interface {
	GetAll(str string) ([]model.Resep, error)
	GetAllByCategoryID(categoryID uint) ([]model.Resep, error)
	GetByID(resepID uint) (model.Resep, error)
	Create(param ResepParam) error
	Update(id uint, param ResepParam) error
}

type ResepParam struct {
	CategoryID uint   `json:"category_id" validate:"required"`
	Judul      string `json:"judul" validate:"required"`
	Deskripsi  string `json:"deskripsi" validate:"required"`
}
