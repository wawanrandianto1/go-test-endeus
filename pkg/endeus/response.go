package endeus

import (
	"endeus/wawan/pkg/model"
	"time"
)

//go:generate mockgen -source=response.go -package=mocks -destination=../../internal/mocks/response.go

type CategoryResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    *[]model.Category `json:"data"`
}

type ResepCustomResponse struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CategoryID uint      `json:"category_id"`
	Judul      string    `json:"judul"`
	Deskripsi  string    `json:"deskripsi"`
	VideoUrl   string    `json:"video_url"`
	Publish    bool      `json:"publish"`
}

type ResepAllResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    *[]ResepCustomResponse `json:"data"`
}

type ResepSingleResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    *model.Resep `json:"data"`
}
