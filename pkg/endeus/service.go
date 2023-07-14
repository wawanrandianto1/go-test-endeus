package endeus

import "endeus/wawan/pkg/model"

//go:generate mockgen -source=service.go -package=mocks -destination=../../internal/mocks/service.go

type ResepService interface {
	GetAll(str string) ([]ResepCustomResponse, error)
	GetAllByCategoryID(catID uint) ([]ResepCustomResponse, error)
	GetByID(id uint) (model.Resep, error)
	CreateNew(param ResepParam) error
	Update(id uint, param ResepParam) error
	Unpublish(id uint) error
	Delete(id uint) error
}

type CategoryService interface {
	GetAll() ([]model.Category, error)
}
