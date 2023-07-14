package model

import "time"

type Resep struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CategoryID uint      `gorm:"column:category_id" json:"category_id"`
	Judul      string    `gorm:"column:judul;not null" json:"judul"`
	Deskripsi  string    `gorm:"column:deskripsi" json:"deskripsi"`
	VideoUrl   string    `gorm:"column:video_url" json:"video_url"`
	Publish    bool      `gorm:"column:publish" json:"publish"`
	Bahan      Bahan     `gorm:"foreignKey:ResepID" json:"bahan"`
	CaraBuat   CaraBuat  `gorm:"foreignKey:ResepID" json:"cara_masak"`
}
