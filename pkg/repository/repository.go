package repository

import (
	"endeus/wawan/pkg/endeus"

	"gorm.io/gorm"
)

type Repo struct {
	Category  endeus.CategoryRepository
	Resep     endeus.ResepRepository
	Bahan     endeus.BahanRepository
	CaraMasak endeus.CaraMasakRepository
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		Resep:     NewResep(db),
		Category:  NewCategory(db),
		CaraMasak: NewMasak(db),
		Bahan:     NewBahan(db),
	}
}
