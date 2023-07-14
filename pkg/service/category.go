package service

import (
	"endeus/wawan/pkg/endeus"
	"endeus/wawan/pkg/model"
	"endeus/wawan/pkg/repository"
)

type categoryService struct {
	*repository.Repo
}

func NewCategoryService(repo *repository.Repo) endeus.CategoryService {
	return &categoryService{
		repo,
	}
}

func (s *categoryService) GetAll() ([]model.Category, error) {
	return s.Category.FindAll()
}
