package model

import "time"

type CaraBuat struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ResepID   uint      `gorm:"column:resep_id;not null" json:"resep_id"`
	LamaWaktu uint      `gorm:"column:lama_waktu" json:"lama_waktu"`
	Deskripsi string    `gorm:"column:deskripsi" json:"deskripsi"`
	Tips      string    `gorm:"column:tips" json:"tips"`
}
