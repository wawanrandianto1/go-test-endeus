package model

import "time"

type Bahan struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ResepID   uint      `gorm:"column:resep_id;not null" json:"resep_id"`
	Deskripsi string    `gorm:"column:deskripsi" json:"deskripsi"`
	Resep     Resep     `gorm:"foreignKey:ResepID"`
}
