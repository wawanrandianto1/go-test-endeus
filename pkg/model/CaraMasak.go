package model

import "time"

type CaraMasak struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ResepID   uint      `gorm:"column:resep_id;not null" json:"resep_id"`
	VideoUrl  string    `gorm:"column:video_url" json:"video_url"`
	LamaMasak uint      `gorm:"column:lama_masak" json:"lama_masak"`
	Deskripsi string    `gorm:"column:deskripsi" json:"deskripsi"`
	Resep     Resep     `gorm:"foreignKey:ResepID"`
}
