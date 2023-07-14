package service

import (
	"endeus/wawan/pkg/endeus"
	"endeus/wawan/pkg/repository"
)

type Service struct {
	endeus.ResepService
	endeus.CategoryService
}

func NewService(repo *repository.Repo) *Service {
	return &Service{
		ResepService:    NewResepService(repo),
		CategoryService: NewCategoryService(repo),
	}
}
